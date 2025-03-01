package in

import (
	"github.com/llamadeus/ebike3/packages/stations/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/stations/domain/port/in"
	"github.com/llamadeus/ebike3/packages/stations/infrastructure/micro"
)

func MakeDeleteStationHandler(stationService in.StationService) micro.HTTPHandler {
	type params struct {
		ID string `param:"id"`
	}

	return micro.MakeHandler(func(ctx micro.Context[params, any]) (*dto.StationDTO, error) {
		id, err := dto.IDFromDTO(ctx.Params().ID)
		if err != nil {
			return nil, micro.NewBadRequestError("invalid station id")
		}

		station, err := stationService.DeleteStation(id)
		if station == nil {
			return nil, err
		}

		return dto.StationToDTO(station), nil
	})
}
