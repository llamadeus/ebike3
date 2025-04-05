package events

import (
	"github.com/guregu/null/v5"
	"time"
)

const (
	RentalsTopic                  = "rentals"
	RentalsRentalStartedEventType = "RentalStarted"
	RentalsRentalStoppedEventType = "RentalStopped"
	RentalsCostUpdatedType        = "CostUpdated"
)

type RentalStartedEvent struct {
	ID          string    `json:"id"`
	CustomerID  string    `json:"customerId"`
	VehicleID   string    `json:"vehicleId"`
	VehicleType string    `json:"vehicleType"`
	Start       time.Time `json:"start"`
	End         null.Time `json:"end"`
}

type RentalStoppedEvent struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customerId"`
	VehicleID  string    `json:"vehicleId"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
}

type CostUpdatedEvent struct {
	ID          string    `json:"id"`
	CustomerID  string    `json:"customerId"`
	VehicleID   string    `json:"vehicleId"`
	VehicleType string    `json:"vehicleType"`
	Start       time.Time `json:"start"`
	Cost        int32     `json:"cost"`
}
