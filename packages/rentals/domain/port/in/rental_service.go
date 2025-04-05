package in

import (
	"context"
	"github.com/llamadeus/ebike3/packages/rentals/domain/model"
	"time"
)

type RentalService interface {
	GetRentalView(id uint64) (*model.RentalView, error)

	GetActiveRentalForCustomer(customerID uint64) (*model.RentalView, error)

	GetPastRentalsForCustomer(customerID uint64) ([]*model.RentalView, error)

	StartRental(ctx context.Context, customerID uint64, vehicleID uint64) (*model.Rental, error)

	StopRental(id uint64, customerID uint64) (*model.Rental, error)

	CreateRentalView(id uint64, customerID uint64, vehicleID uint64, vehicleType model.VehicleType, start time.Time) error

	UpdateRentalView(id uint64, end time.Time) error

	AddExpenseToRental(rentalID uint64, amount int32) error
}
