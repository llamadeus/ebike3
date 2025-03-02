package out

import "github.com/llamadeus/ebike3/packages/vehicles/domain/model"

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

	// UpdateVehicleAvailability updates the availability of the vehicle with the given id.
	UpdateVehicleAvailability(id uint64, available bool) (*model.VehicleView, error)

	// DeleteVehicle deletes the vehicle with the given id.
	DeleteVehicle(id uint64) error
}
