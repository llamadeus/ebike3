package in

import "github.com/llamadeus/ebike3/packages/rentals/domain/model"

type VehicleService interface {
	GetVehicleByID(id uint64) (*model.VehicleView, error)

	CreateVehicleView(id uint64, type_ model.VehicleType) error

	DeleteVehicleView(id uint64) error
}
