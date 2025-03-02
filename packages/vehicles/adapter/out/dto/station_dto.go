package dto

import (
	"fmt"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/model"
	"time"
)

type VehicleDTO struct {
	ID        string  `json:"id"`
	Type      string  `json:"type"`
	PositionX float64 `json:"positionX"`
	PositionY float64 `json:"positionY"`
	Battery   float64 `json:"battery"`
	Available bool    `json:"available"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

func VehicleToDTO(vehicle *model.Vehicle) *VehicleDTO {
	return &VehicleDTO{
		ID:        IDToDTO(vehicle.ID),
		Type:      TypeToDTO(vehicle.Type),
		PositionX: vehicle.PositionX,
		PositionY: vehicle.PositionY,
		Battery:   vehicle.Battery,
		Available: true,
		CreatedAt: vehicle.CreatedAt.Format(time.RFC3339),
		UpdatedAt: vehicle.UpdatedAt.Format(time.RFC3339),
	}
}

func VehicleViewToDTO(vehicle *model.VehicleView) *VehicleDTO {
	return &VehicleDTO{
		ID:        IDToDTO(vehicle.ID),
		Type:      TypeToDTO(vehicle.Type),
		PositionX: vehicle.PositionX,
		PositionY: vehicle.PositionY,
		Battery:   vehicle.Battery,
		Available: vehicle.Available,
		CreatedAt: vehicle.CreatedAt.Format(time.RFC3339),
		UpdatedAt: vehicle.UpdatedAt.Format(time.RFC3339),
	}
}

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
