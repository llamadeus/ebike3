package in

import (
	"github.com/llamadeus/ebike3/packages/vehicles/domain/model"
	"time"
)

type VehicleService interface {
	// CreateVehicle creates a new vehicle with the given type, position, and battery.
	CreateVehicle(type_ model.VehicleType, positionX float64, positionY float64, battery float64) (*model.Vehicle, error)

	// DeleteVehicle deletes the vehicle with the given id.
	DeleteVehicle(id uint64) (*model.Vehicle, error)

	// GetVehicleViews returns all the vehicle views.
	GetVehicleViews() ([]*model.VehicleView, error)

	// GetAvailableVehicleViews returns all the available vehicle views.
	GetAvailableVehicleViews() ([]*model.VehicleView, error)

	// CreateVehicleView creates a new vehicle with the given type, position, and battery.
	CreateVehicleView(id uint64, type_ model.VehicleType, positionX float64, positionY float64, battery float64) error

	// UpdateVehicleView updates the vehicle with the given id.
	UpdateVehicleView(id uint64, positionX float64, positionY float64, battery float64) error

	// UpdateVehicleViewActiveRental updates the active rental of the vehicle with the given id.
	UpdateVehicleViewActiveRental(rentalID uint64, customerID uint64, vehicleID uint64, vehicleType string, start time.Time, cost int32) error

	// ResetVehicleViewActiveRental deletes the active rental of the vehicle with the given id if the rental id matches.
	ResetVehicleViewActiveRental(rentalID uint64, vehicleID uint64) error

	// DeleteVehicleView deletes the vehicle with the given id.
	DeleteVehicleView(id uint64) error
}
