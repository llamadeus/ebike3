package in

import (
	"github.com/llamadeus/ebike3/packages/customers/domain/model"
	"time"
)

type CustomerService interface {
	// UpdateCustomerPosition updates the customer position with the given id.
	UpdateCustomerPosition(id uint64, positionX float64, positionY float64) error

	// GetCustomerViews returns all customers.
	GetCustomerViews() ([]*model.CustomerView, error)

	// GetCustomerViewByID returns the customer with the given id.
	GetCustomerViewByID(id uint64) (*model.CustomerView, error)

	// CreateCustomerView creates a new customer with the given name.
	CreateCustomerView(id uint64, name string) error

	// UpdateCustomerViewPosition updates the customer position with the given id.
	UpdateCustomerViewPosition(id uint64, positionX float64, positionY float64) error

	// UpdateCustomerViewCreditBalance adds the given amount to the customer credit with the given id.
	UpdateCustomerViewCreditBalance(id uint64, amount int32) error

	// UpdateCustomerViewLastLogin updates the customer last login with the given id.
	UpdateCustomerViewLastLogin(id uint64, lastLogin time.Time) error

	// UpdateCustomerViewActiveRental updates the active rental of the customer with the given id.
	UpdateCustomerViewActiveRental(id uint64, rentalID uint64, vehicleID uint64, vehicleType string, start time.Time) error

	// ResetCustomerViewActiveRental deletes the active rental of the customer with the given id if the rental id matches.
	ResetCustomerViewActiveRental(id uint64, rentalID uint64) error
}
