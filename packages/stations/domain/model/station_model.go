package model

import (
	"time"
)

type Station struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	PositionX float64   `db:"position_x"`
	PositionY float64   `db:"position_y"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type StationView struct {
	ID        uint64    `bson:"_id,omitempty"`
	Name      string    `bson:"name,omitempty"`
	PositionX float64   `bson:"positionX,omitempty"`
	PositionY float64   `bson:"positionY,omitempty"`
	CreatedAt time.Time `bson:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty"`
}
