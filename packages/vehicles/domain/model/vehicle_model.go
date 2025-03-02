package model

import (
	"time"
)

type Vehicle struct {
	ID        uint64      `db:"id"`
	Type      VehicleType `db:"type"`
	PositionX float64     `db:"position_x"`
	PositionY float64     `db:"position_y"`
	Battery   float64     `db:"battery"`
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt time.Time   `db:"updated_at"`
}

type VehicleView struct {
	ID        uint64      `bson:"_id,omitempty"`
	Type      VehicleType `bson:"type,omitempty"`
	PositionX float64     `bson:"positionX,omitempty"`
	PositionY float64     `bson:"positionY,omitempty"`
	Battery   float64     `bson:"battery,omitempty"`
	Available bool        `bson:"available,omitempty"`
	CreatedAt time.Time   `bson:"createdAt,omitempty"`
	UpdatedAt time.Time   `bson:"updatedAt,omitempty"`
}

type VehicleType string

const (
	VehicleTypeBike  VehicleType = "BIKE"
	VehicleTypeEBike VehicleType = "EBIKE"
	VehicleTypeABike VehicleType = "ABIKE"
)
