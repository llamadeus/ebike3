package out

import (
	"github.com/llamadeus/ebike3/packages/rentals/domain/model"
	"time"
)

type RentalViewRepository interface {
	Get(id uint64) (*model.RentalView, error)

	GetActiveRentalByCustomerID(customerID uint64) (*model.RentalView, error)

	GetPastRentalsByCustomerID(customerID uint64) ([]*model.RentalView, error)

	Create(id uint64, customerID uint64, vehicleID uint64, vehicleType model.VehicleType, start time.Time) (*model.RentalView, error)

	Update(id uint64, end time.Time) (*model.RentalView, error)

	AddExpense(rentalID uint64, amount int32) (*model.RentalView, error)
}
