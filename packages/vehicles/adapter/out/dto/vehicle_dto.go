package dto

import (
	"fmt"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/model"
	"time"
)

type VehicleDTO struct {
	ID           string     `json:"id"`
	Type         string     `json:"type"`
	PositionX    float64    `json:"positionX"`
	PositionY    float64    `json:"positionY"`
	Battery      float64    `json:"battery"`
	Available    bool       `json:"available"`
	ActiveRental *RentalDTO `json:"activeRental"`
	CreatedAt    string     `json:"createdAt"`
	UpdatedAt    string     `json:"updatedAt"`
}

type RentalDTO struct {
	ID          string `json:"id"`
	VehicleID   string `json:"vehicleId"`
	CustomerID  string `json:"customerId"`
	VehicleType string `json:"vehicleType"`
	Start       string `json:"start"`
	Cost        int32  `json:"cost"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
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
		ID:           IDToDTO(vehicle.ID),
		Type:         TypeToDTO(vehicle.Type),
		PositionX:    vehicle.PositionX,
		PositionY:    vehicle.PositionY,
		Battery:      vehicle.Battery,
		Available:    vehicle.ActiveRental == nil,
		ActiveRental: RentalViewToDTO(vehicle.ActiveRental),
		CreatedAt:    vehicle.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    vehicle.UpdatedAt.Format(time.RFC3339),
	}
}

func RentalViewToDTO(rental *model.RentalView) *RentalDTO {
	if rental == nil {
		return nil
	}

	return &RentalDTO{
		ID:          IDToDTO(rental.ID),
		VehicleID:   IDToDTO(rental.VehicleID),
		CustomerID:  IDToDTO(rental.CustomerID),
		VehicleType: rental.VehicleType,
		Start:       rental.Start.Format(time.RFC3339),
		Cost:        rental.Cost,
		CreatedAt:   rental.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   rental.UpdatedAt.Format(time.RFC3339),
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
