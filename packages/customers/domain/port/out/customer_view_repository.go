package out

import (
	"github.com/llamadeus/ebike3/packages/customers/domain/model"
	"time"
)

// CustomerViewRepository is an interface for a Mongo repository for customer view data.
type CustomerViewRepository interface {
	// GetCustomers returns all the customers.
	GetCustomers() ([]*model.CustomerView, error)

	// GetCustomerByID returns the customer with the given id.
	GetCustomerByID(id uint64) (*model.CustomerView, error)

	// CreateCustomer creates a new customer with the given name and position.
	CreateCustomer(id uint64, name string, positionX float64, positionY float64, creditBalance int32) error

	// UpdateCustomerViewPosition updates the customer position with the given id.
	UpdateCustomerViewPosition(id uint64, positionX float64, positionY float64) error

	// UpdateCustomerViewCreditBalance updates the customer credit with the given id.
	UpdateCustomerViewCreditBalance(id uint64, creditBalance int32) error

	// UpdateCustomerViewLastLogin updates the customer last login with the given id.
	UpdateCustomerViewLastLogin(id uint64, lastLogin time.Time) error

	// UpdateCustomerViewActiveRental updates the active rental of the customer with the given id.
	UpdateCustomerViewActiveRental(customerID uint64, rentalID uint64, vehicleID uint64, vehicleType string, start time.Time) error
}
