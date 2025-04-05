package out

import (
	"github.com/llamadeus/ebike3/packages/vehicles/domain/model"
	"time"
)

// VehicleViewRepository is an interface for a Mongo repository for vehicle view data.
type VehicleViewRepository interface {
	// GetVehicles returns all the vehicles.
	GetVehicles() ([]*model.VehicleView, error)

	// GetAvailableVehicles returns all the available vehicles.
	GetAvailableVehicles() ([]*model.VehicleView, error)

	// GetVehicleByID returns the vehicle with the given id.
	GetVehicleByID(id uint64) (*model.VehicleView, error)

	// CreateVehicle creates a new vehicle with the given type, position, and battery.
	CreateVehicle(id uint64, type_ model.VehicleType, positionX float64, positionY float64, battery float64) (*model.VehicleView, error)

	// UpdateVehicle updates the vehicle with the given id.
	UpdateVehicle(id uint64, positionX float64, positionY float64, battery float64) (*model.VehicleView, error)

	// UpdateActiveRental updates the active rental of the vehicle with the given id.
	UpdateActiveRental(rentalID uint64, customerID uint64, vehicleID uint64, vehicleType string, start time.Time, cost int32) error

	// ResetActiveRental deletes the active rental of the vehicle with the given id.
	ResetActiveRental(id uint64) error

	// DeleteVehicle deletes the vehicle with the given id.
	DeleteVehicle(id uint64) error
}
