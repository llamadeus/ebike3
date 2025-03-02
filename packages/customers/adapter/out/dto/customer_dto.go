package dto

import (
	"github.com/guregu/null/v5"
	"github.com/llamadeus/ebike3/packages/customers/domain/model"
	"time"
)

type CustomerDTO struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	PositionX     float64     `json:"positionX"`
	PositionY     float64     `json:"positionY"`
	CreditBalance float64     `json:"creditBalance"`
	ActiveRental  *RentalDTO  `json:"activeRental"`
	LastLogin     null.String `json:"lastLogin"`
	CreatedAt     string      `json:"createdAt"`
	UpdatedAt     string      `json:"updatedAt"`
}

type RentalDTO struct {
	ID          string `json:"id"`
	VehicleID   string `json:"vehicleId"`
	CustomerID  string `json:"customerId"`
	VehicleType string `json:"vehicleType"`
	Start       string `json:"start"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func CustomerViewToDTO(customer *model.CustomerView) *CustomerDTO {
	var lastLogin null.String
	if customer.LastLogin.Valid {
		lastLogin.SetValid(customer.LastLogin.Time.Format(time.RFC3339))
	}

	return &CustomerDTO{
		ID:            IDToDTO(customer.ID),
		Name:          customer.Name,
		PositionX:     customer.PositionX,
		PositionY:     customer.PositionY,
		CreditBalance: customer.CreditBalance,
		ActiveRental:  RentalViewToDTO(customer.ActiveRental),
		LastLogin:     lastLogin,
		CreatedAt:     customer.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     customer.UpdatedAt.Format(time.RFC3339),
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
		CreatedAt:   rental.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   rental.UpdatedAt.Format(time.RFC3339),
	}
}
