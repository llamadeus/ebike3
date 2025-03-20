package out

import "github.com/llamadeus/ebike3/packages/rentals/domain/model"

type RentalRepository interface {
	Get(id uint64) (*model.Rental, error)

	GetActiveRentalByCustomerID(customerID uint64) (*model.Rental, error)

	GetPastRentalsByCustomerID(customerID uint64) ([]*model.Rental, error)

	GetActiveRentalByVehicleID(vehicleID uint64) (*model.Rental, error)

	CreateRental(customerID uint64, vehicleID uint64) (*model.Rental, error)

	StopRental(id uint64) (*model.Rental, error)
}
