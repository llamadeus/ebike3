package model

import (
	"github.com/guregu/null/v5"
	"time"
)

type Rental struct {
	ID         uint64    `db:"id"`
	CustomerID uint64    `db:"customer_id"`
	VehicleID  uint64    `db:"vehicle_id"`
	Start      time.Time `db:"start"`
	End        null.Time `db:"end"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type RentalView struct {
	ID          uint64      `bson:"_id,omitempty"`
	CustomerID  uint64      `bson:"customerId,omitempty"`
	VehicleID   uint64      `bson:"vehicleId,omitempty"`
	VehicleType VehicleType `bson:"vehicleType,omitempty"`
	Start       time.Time   `bson:"start,omitempty"`
	End         null.Time   `bson:"end,omitempty"`
	Cost        int32       `bson:"cost,omitempty"`
	CreatedAt   time.Time   `bson:"createdAt,omitempty"`
	UpdatedAt   time.Time   `bson:"updatedAt,omitempty"`
}
