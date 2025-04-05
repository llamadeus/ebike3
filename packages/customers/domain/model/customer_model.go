package model

import (
	"github.com/guregu/null/v5"
	"time"
)

type CustomerView struct {
	ID            uint64      `bson:"_id,omitempty"`
	Name          string      `bson:"name,omitempty"`
	PositionX     float64     `bson:"positionX,omitempty"`
	PositionY     float64     `bson:"positionY,omitempty"`
	CreditBalance int32       `bson:"creditBalance,omitempty"`
	ActiveRental  *RentalView `bson:"activeRental,omitempty"`
	LastLogin     null.Time   `bson:"lastLogin,omitempty"`
	CreatedAt     time.Time   `bson:"createdAt,omitempty"`
	UpdatedAt     time.Time   `bson:"updatedAt,omitempty"`
}

type RentalView struct {
	ID          uint64    `bson:"_id,omitempty"`
	VehicleID   uint64    `bson:"vehicleId,omitempty"`
	CustomerID  uint64    `bson:"customerId,omitempty"`
	VehicleType string    `bson:"vehicleType,omitempty"`
	Start       time.Time `bson:"start,omitempty"`
	Cost        int32     `bson:"cost,omitempty"`
	CreatedAt   time.Time `bson:"createdAt,omitempty"`
	UpdatedAt   time.Time `bson:"updatedAt,omitempty"`
}
