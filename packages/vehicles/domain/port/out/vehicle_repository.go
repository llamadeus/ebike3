package out

import "github.com/llamadeus/ebike3/packages/vehicles/domain/model"

// VehicleRepository is an interface for a repository that handles vehicle operations.
type VehicleRepository interface {
	// GetVehicles returns all the vehicles.
	GetVehicles() ([]*model.Vehicle, error)

	// GetVehicleByID returns the vehicle with the given id.
	GetVehicleByID(id uint64) (*model.Vehicle, error)

	// CreateVehicle creates a new vehicle with the given type, position, and battery.
	CreateVehicle(type_ model.VehicleType, positionX float64, positionY float64, battery float64) (*model.Vehicle, error)

	// UpdateVehicle updates the vehicle with the given id.
	UpdateVehicle(id uint64, positionX float64, positionY float64, battery float64) (*model.Vehicle, error)

	// DeleteVehicle deletes the vehicle with the given id.
	DeleteVehicle(id uint64) error
}
