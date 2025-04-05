package service

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/llamadeus/ebike3/packages/rentals/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/rentals/domain/constants"
	"github.com/llamadeus/ebike3/packages/rentals/domain/events"
	"github.com/llamadeus/ebike3/packages/rentals/domain/model"
	"github.com/llamadeus/ebike3/packages/rentals/domain/port/in"
	"github.com/llamadeus/ebike3/packages/rentals/domain/port/out"
	"github.com/llamadeus/ebike3/packages/rentals/domain/tasks"
	"github.com/llamadeus/ebike3/packages/rentals/infrastructure/micro"
	"log/slog"
	"time"
)

type RentalService struct {
	kafka          micro.Kafka
	asynq          *asynq.Client
	repository     out.RentalRepository
	viewRepository out.RentalViewRepository
	vehicleService in.VehicleService
}

var _ in.RentalService = (*RentalService)(nil)

func NewRentalService(kafka micro.Kafka, asynq *asynq.Client, repository out.RentalRepository, viewRepository out.RentalViewRepository, vehicleService in.VehicleService) *RentalService {
	return &RentalService{
		kafka:          kafka,
		asynq:          asynq,
		repository:     repository,
		viewRepository: viewRepository,
		vehicleService: vehicleService,
	}
}

func (s *RentalService) GetRentalView(id uint64) (*model.RentalView, error) {
	return s.viewRepository.Get(id)
}

func (s *RentalService) GetActiveRentalForCustomer(customerID uint64) (*model.RentalView, error) {
	rental, err := s.viewRepository.GetActiveRentalByCustomerID(customerID)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to get active rental: %v", err))
	}

	return rental, nil
}

func (s *RentalService) GetPastRentalsForCustomer(customerID uint64) ([]*model.RentalView, error) {
	rentals, err := s.viewRepository.GetPastRentalsByCustomerID(customerID)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to get past rentals: %v", err))
	}

	return rentals, nil
}

func (s *RentalService) StartRental(ctx context.Context, customerID uint64, vehicleID uint64) (*model.Rental, error) {
	vehicle, err := s.vehicleService.GetVehicleByID(vehicleID)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to get vehicle: %v", err))
	}

	if vehicle == nil {
		return nil, micro.NewNotFoundError(fmt.Sprintf("vehicle with id %d not found", vehicleID))
	}

	activeRental, err := s.repository.GetActiveRentalByCustomerID(customerID)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to get active rental: %v", err))
	}

	if activeRental != nil {
		return nil, micro.NewBadRequestError("customer already has an active rental")
	}

	activeRental, err = s.repository.GetActiveRentalByVehicleID(vehicleID)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to get active rental: %v", err))
	}

	if activeRental != nil {
		return nil, micro.NewBadRequestError("vehicle already rented")
	}

	// Check if the customer's credit balance is sufficient
	fee := s.getUnblockingFee(vehicle.Type)
	preliminaryExpenseID, err := s.createPreliminaryExpense(ctx, customerID, fee)
	if err != nil {
		return nil, err
	}

	rental, err := s.repository.CreateRental(customerID, vehicleID)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to create rental: %v", err))
	}

	_, err = s.viewRepository.Create(rental.ID, customerID, vehicleID, vehicle.Type, rental.Start)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to create rental view: %v", err))
	}

	err = s.finalizePreliminaryExpense(ctx, preliminaryExpenseID, rental.ID)
	if err != nil {
		// TODO: Maybe we should delete the rental right here

		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to finalize preliminary expense: %v", err))
	}

	task, err := tasks.NewRentalsChargeActiveRentalTask(dto.IDToDTO(rental.ID))
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to create charge active rental task: %v", err))
	}
	_, err = s.asynq.Enqueue(task, asynq.ProcessIn(60*time.Second))
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to enqueue charge active rental task: %v", err))
	}

	event := micro.NewEvent(events.RentalsRentalStartedEventType, events.RentalStartedEvent{
		ID:          dto.IDToDTO(rental.ID),
		CustomerID:  dto.IDToDTO(rental.CustomerID),
		VehicleID:   dto.IDToDTO(rental.VehicleID),
		VehicleType: dto.TypeToDTO(vehicle.Type),
		Start:       rental.Start,
		End:         rental.End,
	})
	err = s.kafka.Producer().Send(events.RentalsTopic, event.Payload.ID, event)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to send kafka event: %v", err))
	}

	return rental, nil
}

func (s *RentalService) StopRental(id uint64, customerID uint64) (*model.Rental, error) {
	rental, err := s.repository.Get(id)
	if err != nil {
		return nil, micro.NewNotFoundError(fmt.Sprintf("rental with id %d not found", id))
	}

	if rental.CustomerID != customerID {
		return nil, micro.NewUnauthorizedError(fmt.Sprintf("customer %d not authorized to stop rental %d", customerID, id))
	}

	stopped, err := s.repository.StopRental(id)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to stop rental: %v", err))
	}

	event := micro.NewEvent(events.RentalsRentalStoppedEventType, events.RentalStoppedEvent{
		ID:         dto.IDToDTO(stopped.ID),
		CustomerID: dto.IDToDTO(stopped.CustomerID),
		VehicleID:  dto.IDToDTO(stopped.VehicleID),
		Start:      stopped.Start,
		End:        stopped.End.Time,
	})
	err = s.kafka.Producer().Send(events.RentalsTopic, event.Payload.ID, event)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to send kafka event: %v", err))
	}

	return stopped, nil
}

func (s *RentalService) CreateRentalView(id uint64, customerID uint64, vehicleID uint64, vehicleType model.VehicleType, start time.Time) error {
	_, err := s.viewRepository.Create(id, customerID, vehicleID, vehicleType, start)

	return err
}

func (s *RentalService) UpdateRentalView(id uint64, end time.Time) error {
	_, err := s.viewRepository.Update(id, end)

	return err
}

func (s *RentalService) AddExpenseToRental(rentalID uint64, amount int32) error {
	rental, err := s.viewRepository.AddExpense(rentalID, amount)
	if err != nil {
		return err
	}

	event := micro.NewEvent(events.RentalsCostUpdatedType, events.CostUpdatedEvent{
		ID:          dto.IDToDTO(rental.ID),
		CustomerID:  dto.IDToDTO(rental.CustomerID),
		VehicleID:   dto.IDToDTO(rental.VehicleID),
		VehicleType: dto.TypeToDTO(rental.VehicleType),
		Start:       rental.Start,
		Cost:        rental.Cost,
	})
	err = s.kafka.Producer().Send(events.RentalsTopic, event.Payload.ID, event)
	if err != nil {
		return err
	}

	return nil
}

func (s *RentalService) ChargeActiveRental(ctx context.Context, rentalID uint64) error {
	start := time.Now()
	rental, err := s.repository.Get(rentalID)
	if err != nil {
		return fmt.Errorf("failed to get rental: %v", err)
	}
	if rental == nil {
		return fmt.Errorf("rental with id %d not found", rentalID)
	}

	if rental.End.Valid {
		// Rental has already ended, nothing to do
		return nil
	}

	vehicle, err := s.vehicleService.GetVehicleByID(rental.VehicleID)
	if err != nil {
		return fmt.Errorf("failed to get vehicle: %v", err)
	}
	if vehicle == nil {
		return fmt.Errorf("vehicle with id %d not found", rental.VehicleID)
	}

	// Create a new expense for the rental (PUT /expenses to accounting service)
	err = s.createExpense(ctx, rental.CustomerID, rental.ID, s.getRentalFeePerMinute(vehicle.Type))
	if err != nil {
		return err
	}

	// Queue a new task to charge the rental in (60 - delta) seconds, where `delta = time.Now() - start`
	delta := time.Now().Sub(start)
	nextCharge := time.Duration(60-delta.Seconds()) * time.Second
	slog.Info(
		"queueing task to charge rental",
		"rentalId", rental.ID,
		"nextCharge", nextCharge,
	)

	task, err := tasks.NewRentalsChargeActiveRentalTask(dto.IDToDTO(rental.ID))
	if err != nil {
		return err
	}
	_, err = s.asynq.Enqueue(task, asynq.ProcessIn(nextCharge))
	if err != nil {
		return err
	}

	return nil
}

func (s *RentalService) getUnblockingFee(vehicleType model.VehicleType) int32 {
	switch vehicleType {
	case model.VehicleTypeBike:
		return constants.UnblockingFeeBike
	case model.VehicleTypeEBike:
		return constants.UnblockingFeeEBike
	case model.VehicleTypeABike:
		return constants.UnblockingFeeABike
	}

	return 0
}

func (s *RentalService) getRentalFeePerMinute(vehicleType model.VehicleType) int32 {
	switch vehicleType {
	case model.VehicleTypeBike:
		return constants.RentalFeePerMinuteBike
	case model.VehicleTypeEBike:
		return constants.RentalFeePerMinuteEBike
	case model.VehicleTypeABike:
		return constants.RentalFeePerMinuteABike
	}

	return 0
}

func (s *RentalService) createExpense(ctx context.Context, customerID uint64, rentalID uint64, amount int32) error {
	type expenseInput struct {
		CustomerID string `json:"customerId"`
		RentalID   string `json:"rentalId"`
		Amount     int32  `json:"amount"`
	}

	var err error

	for i := 0; i < 5; i++ {
		_, err = micro.Invoke[expenseInput, any](ctx, "PUT accounting-service:5001/expenses", nil, expenseInput{
			CustomerID: dto.IDToDTO(customerID),
			RentalID:   dto.IDToDTO(rentalID),
			Amount:     amount,
		})
		if err == nil {
			return nil
		}

		// Sleep for an increasing amount of time before retrying
		time.Sleep(time.Second * time.Duration(i+1))
	}

	return err
}

func (s *RentalService) createPreliminaryExpense(ctx context.Context, customerID uint64, amount int32) (string, error) {
	type preliminaryExpenseInput struct {
		InquiryID  string `json:"inquiryId"`
		CustomerID string `json:"customerId"`
		Amount     int32  `json:"amount"`
	}

	type preliminaryExpenseDTO struct {
		ID         string `json:"id"`
		InquiryID  string `json:"inquiryId"`
		CustomerID string `json:"customerId"`
		Amount     int32  `json:"amount"`
		CreatedAt  string `json:"createdAt"`
		ExpiresAt  string `json:"expiresAt"`
	}

	var inquiryID uint64
	if err := binary.Read(rand.Reader, binary.BigEndian, &inquiryID); err != nil {
		return "", err
	}

	// Set highest bit of inquiry id to 0 to prevent Postgres issues
	inquiryID = inquiryID &^ (1 << 63)

	slog.Info(
		"creating preliminary expense",
		"inquiryId", inquiryID,
		"customerId", customerID,
		"amount", amount,
	)

	var preliminaryExpense preliminaryExpenseDTO
	var err error

	for i := 0; i < 5; i++ {
		preliminaryExpense, err = micro.Invoke[preliminaryExpenseInput, preliminaryExpenseDTO](ctx, "PUT accounting-service:5001/preliminary-expenses", nil, preliminaryExpenseInput{
			InquiryID:  dto.IDToDTO(inquiryID),
			CustomerID: dto.IDToDTO(customerID),
			Amount:     amount,
		})
		if err == nil {
			return preliminaryExpense.ID, nil
		}

		var invokeError *micro.InvokeError
		if errors.As(err, &invokeError) {
			if invokeError.Status == 400 {
				return "", fmt.Errorf("customer %d does not have enough credit balance", customerID)
			}
		}

		// Sleep for an increasing amount of time before retrying
		time.Sleep(time.Second * time.Duration(i+1))
	}

	return "", err
}

func (s *RentalService) finalizePreliminaryExpense(ctx context.Context, id string, rentalID uint64) error {
	type input struct {
		RentalID string `json:"rentalId" validate:"required"`
	}

	var err error

	endpoint := fmt.Sprintf("POST accounting-service:5001/preliminary-expenses/%s/finalize", id)
	for i := 0; i < 5; i++ {
		_, err = micro.Invoke[input, any](ctx, endpoint, nil, input{
			RentalID: dto.IDToDTO(rentalID),
		})
		if err == nil {
			return nil
		}

		var invokeError *micro.InvokeError
		if errors.As(err, &invokeError) {
			if invokeError.Status == 404 {
				return nil
			}
		}

		// Sleep for an increasing amount of time before retrying
		time.Sleep(time.Second * time.Duration(i+1))
	}

	return err
}
