package in

import "github.com/llamadeus/ebike3/packages/vehicles/domain/model"

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

	// UpdateVehicleViewAvailability updates the availability of the vehicle with the given id.
	UpdateVehicleViewAvailability(id uint64, available bool) error

	// DeleteVehicleView deletes the vehicle with the given id.
	DeleteVehicleView(id uint64) error
}
