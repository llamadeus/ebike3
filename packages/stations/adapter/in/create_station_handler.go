package in

import (
	"github.com/llamadeus/ebike3/packages/stations/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/stations/domain/port/in"
	"github.com/llamadeus/ebike3/packages/stations/infrastructure/micro"
)

func MakeCreateStationHandler(stationService in.StationService) micro.HTTPHandler {
	type input struct {
		Name      string  `json:"name" validate:"required"`
		PositionX float64 `json:"positionX" validate:"required"`
		PositionY float64 `json:"positionY" validate:"required"`
	}

	return micro.MakeHandler(func(ctx micro.Context[any, input]) (*dto.StationDTO, error) {
		station, err := stationService.CreateStation(ctx.Input().Name, ctx.Input().PositionX, ctx.Input().PositionY)
		if station == nil {
			return nil, err
		}

		return dto.StationToDTO(station), nil
	})
}
