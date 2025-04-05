package service

import (
	"fmt"
	"github.com/llamadeus/ebike3/packages/customers/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/customers/domain/events"
	"github.com/llamadeus/ebike3/packages/customers/domain/model"
	"github.com/llamadeus/ebike3/packages/customers/domain/port/in"
	"github.com/llamadeus/ebike3/packages/customers/domain/port/out"
	"github.com/llamadeus/ebike3/packages/customers/infrastructure/micro"
	"time"
)

// CustomerService implements the CustomerService interface.
type CustomerService struct {
	kafka          micro.Kafka
	viewRepository out.CustomerViewRepository
}

// Ensure that CustomerService implements the CustomerService interface.
var _ in.CustomerService = (*CustomerService)(nil)

// NewCustomerService creates a new instance of the CustomerService.
func NewCustomerService(kafka micro.Kafka, viewRepository out.CustomerViewRepository) *CustomerService {
	return &CustomerService{
		kafka:          kafka,
		viewRepository: viewRepository,
	}
}

// UpdateCustomerPosition updates the customer position with the given id.
func (s *CustomerService) UpdateCustomerPosition(id uint64, positionX float64, positionY float64) error {
	customer, err := s.viewRepository.GetCustomerByID(id)
	if err != nil {
		return micro.NewNotFoundError(fmt.Sprintf("customer with id %d not found", id))
	}

	event := micro.NewEvent(events.CustomersCustomerPositionUpdatedEventType, events.CustomerPositionUpdatedEvent{
		ID:        dto.IDToDTO(customer.ID),
		PositionX: positionX,
		PositionY: positionY,
	})
	err = s.kafka.Producer().Send(events.CustomersTopic, event.Payload.ID, event)
	if err != nil {
		return micro.NewInternalServerError(fmt.Sprintf("failed to send kafka event: %v", err))
	}

	return nil
}

// GetCustomerViews returns all the stations.
func (s *CustomerService) GetCustomerViews() ([]*model.CustomerView, error) {
	return s.viewRepository.GetCustomers()
}

func (s *CustomerService) GetCustomerViewByID(id uint64) (*model.CustomerView, error) {
	return s.viewRepository.GetCustomerByID(id)
}

// CreateCustomerView creates a new station with the given name and position.
func (s *CustomerService) CreateCustomerView(id uint64, name string) error {
	return s.viewRepository.CreateCustomer(id, name, 0, 0, 0)
}

func (s *CustomerService) UpdateCustomerViewPosition(id uint64, positionX float64, positionY float64) error {
	return s.viewRepository.UpdateCustomerViewPosition(id, positionX, positionY)
}

func (s *CustomerService) UpdateCustomerViewCreditBalance(id uint64, amount int32) error {
	return s.viewRepository.UpdateCustomerViewCreditBalance(id, amount)
}

func (s *CustomerService) UpdateCustomerViewLastLogin(id uint64, lastLogin time.Time) error {
	return s.viewRepository.UpdateCustomerViewLastLogin(id, lastLogin)
}

func (s *CustomerService) UpdateCustomerViewActiveRental(id uint64, rentalID uint64, vehicleID uint64, vehicleType string, start time.Time) error {
	return s.viewRepository.UpdateCustomerViewActiveRental(id, rentalID, vehicleID, vehicleType, start)
}

func (s *CustomerService) ResetCustomerViewActiveRental(id uint64, rentalID uint64) error {
	customer, err := s.viewRepository.GetCustomerByID(id)
	if err != nil {
		return err
	}

	if customer.ActiveRental == nil {
		return nil
	}

	if customer.ActiveRental.ID != rentalID {
		return nil
	}

	return s.viewRepository.ResetCustomerViewActiveRental(id)
}
