package in

import (
	"github.com/llamadeus/ebike3/packages/stations/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/stations/domain/port/in"
	"github.com/llamadeus/ebike3/packages/stations/infrastructure/micro"
)

func MakeUpdateStationHandler(stationService in.StationService) micro.HTTPHandler {
	type params struct {
		ID string `param:"id"`
	}

	type input struct {
		Name      string  `json:"name" validate:"required"`
		PositionX float64 `json:"positionX" validate:"required"`
		PositionY float64 `json:"positionY" validate:"required"`
	}

	return micro.MakeHandler(func(ctx micro.Context[params, input]) (*dto.StationDTO, error) {
		id, err := dto.IDFromDTO(ctx.Params().ID)
		if err != nil {
			return nil, micro.NewBadRequestError("invalid station id")
		}

		station, err := stationService.UpdateStation(id, ctx.Input().Name, ctx.Input().PositionX, ctx.Input().PositionY)
		if station == nil {
			return nil, err
		}

		return dto.StationToDTO(station), nil
	})
}
