package service

import (
	"context"
	"fmt"
	"github.com/llamadeus/ebike3/packages/rentals/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/rentals/domain/constants"
	"github.com/llamadeus/ebike3/packages/rentals/domain/events"
	"github.com/llamadeus/ebike3/packages/rentals/domain/model"
	"github.com/llamadeus/ebike3/packages/rentals/domain/port/in"
	"github.com/llamadeus/ebike3/packages/rentals/domain/port/out"
	"github.com/llamadeus/ebike3/packages/rentals/infrastructure/micro"
	"time"
)

type RentalService struct {
	kafka          micro.Kafka
	repository     out.RentalRepository
	viewRepository out.RentalViewRepository
	vehicleService in.VehicleService
}

var _ in.RentalService = (*RentalService)(nil)

func NewRentalService(kafka micro.Kafka, repository out.RentalRepository, viewRepository out.RentalViewRepository, vehicleService in.VehicleService) *RentalService {
	return &RentalService{
		kafka:          kafka,
		repository:     repository,
		viewRepository: viewRepository,
		vehicleService: vehicleService,
	}
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
	type creditBalanceDTO struct {
		CustomerID    string `json:"customerId"`
		CreditBalance int32  `json:"creditBalance"`
	}

	endpoint := fmt.Sprintf("GET accounting-service:5001/customers/%s/credit-balance", dto.IDToDTO(customerID))
	result, err := micro.Invoke[any, creditBalanceDTO](ctx, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	fee := s.getUnblockingFee(vehicle.Type)
	if result.CreditBalance < fee {
		return nil, micro.NewBadRequestError(fmt.Sprintf("customer %d does not have enough credit balance", customerID))
	}

	rental, err := s.repository.CreateRental(customerID, vehicleID)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to create rental: %v", err))
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

	return nil, nil
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
	return s.viewRepository.AddExpense(rentalID, amount)
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
