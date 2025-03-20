package dto

import (
	"fmt"
	"github.com/llamadeus/ebike3/packages/rentals/domain/model"
)

func TypeToDTO(type_ model.VehicleType) string {
	return string(type_)
}

func TypeFromDTO(type_ string) (model.VehicleType, error) {
	switch type_ {
	case string(model.VehicleTypeBike):
		return model.VehicleTypeBike, nil
	case string(model.VehicleTypeEBike):
		return model.VehicleTypeEBike, nil
	case string(model.VehicleTypeABike):
		return model.VehicleTypeABike, nil
	}

	return model.VehicleTypeBike, fmt.Errorf("invalid vehicle type: %s", type_)
}
