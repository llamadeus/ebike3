package model

import (
	"time"
)

type VehicleView struct {
	ID        uint64      `bson:"_id,omitempty"`
	Type      VehicleType `bson:"type,omitempty"`
	CreatedAt time.Time   `bson:"createdAt,omitempty"`
	UpdatedAt time.Time   `bson:"updatedAt,omitempty"`
}

type VehicleType string

const (
	VehicleTypeBike  VehicleType = "BIKE"
	VehicleTypeEBike VehicleType = "EBIKE"
	VehicleTypeABike VehicleType = "ABIKE"
)
