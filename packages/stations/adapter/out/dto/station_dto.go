package dto

import (
	"github.com/llamadeus/ebike3/packages/stations/domain/model"
	"time"
)

type StationDTO struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	PositionX float64 `json:"positionX"`
	PositionY float64 `json:"positionY"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
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
