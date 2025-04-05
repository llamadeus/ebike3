package dto

import (
	"github.com/guregu/null/v5"
	"github.com/llamadeus/ebike3/packages/rentals/domain/model"
	"time"
)

type RentalDTO struct {
	ID         string      `json:"id"`
	CustomerID string      `json:"customerId"`
	VehicleID  string      `json:"vehicleId"`
	Start      string      `json:"start"`
	End        null.String `json:"end"`
	Cost       int32       `json:"cost"`
	CreatedAt  string      `json:"createdAt"`
	UpdatedAt  string      `json:"updatedAt"`
}

func RentalToDTO(rental *model.Rental, cost int32) *RentalDTO {
	var end null.String
	if rental.End.Valid {
		end.SetValid(rental.End.Time.Format(time.RFC3339))
	}

	return &RentalDTO{
		ID:         IDToDTO(rental.ID),
		CustomerID: IDToDTO(rental.CustomerID),
		VehicleID:  IDToDTO(rental.VehicleID),
		Start:      rental.Start.Format(time.RFC3339),
		End:        end,
		Cost:       cost,
		CreatedAt:  rental.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  rental.UpdatedAt.Format(time.RFC3339),
	}
}

func RentalViewToDTO(rental *model.RentalView) *RentalDTO {
	var end null.String
	if rental.End.Valid {
		end.SetValid(rental.End.Time.Format(time.RFC3339))
	}

	return &RentalDTO{
		ID:         IDToDTO(rental.ID),
		CustomerID: IDToDTO(rental.CustomerID),
		VehicleID:  IDToDTO(rental.VehicleID),
		Start:      rental.Start.Format(time.RFC3339),
		End:        end,
		Cost:       rental.Cost,
		CreatedAt:  rental.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  rental.UpdatedAt.Format(time.RFC3339),
	}
}
