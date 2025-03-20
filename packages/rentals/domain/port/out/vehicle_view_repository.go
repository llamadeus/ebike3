package out

import (
	"github.com/llamadeus/ebike3/packages/rentals/domain/model"
)

type VehicleViewRepository interface {
	Get(id uint64) (*model.VehicleView, error)

	Create(id uint64, type_ model.VehicleType) (*model.VehicleView, error)

	Delete(id uint64) error
}
