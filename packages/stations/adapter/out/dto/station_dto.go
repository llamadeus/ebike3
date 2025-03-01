package dto

import (
	"github.com/llamadeus/ebike3/packages/stations/domain/model"
	"time"
)

type StationDTO struct {
	ID        string  `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	PositionX float64 `json:"positionX,omitempty"`
	PositionY float64 `json:"positionY,omitempty"`
	CreatedAt string  `json:"createdAt,omitempty"`
	UpdatedAt string  `json:"updatedAt,omitempty"`
}

func StationToDTO(station *model.Station) *StationDTO {
	return &StationDTO{
		ID:        IDToDTO(station.ID),
		Name:      station.Name,
		PositionX: station.PositionX,
		PositionY: station.PositionY,
		CreatedAt: station.CreatedAt.Format(time.RFC3339),
		UpdatedAt: station.UpdatedAt.Format(time.RFC3339),
	}
}

func StationViewToDTO(station *model.StationView) *StationDTO {
	return &StationDTO{
		ID:        IDToDTO(station.ID),
		Name:      station.Name,
		PositionX: station.PositionX,
		PositionY: station.PositionY,
		CreatedAt: station.CreatedAt.Format(time.RFC3339),
		UpdatedAt: station.UpdatedAt.Format(time.RFC3339),
	}
}
