package in

import (
	"github.com/llamadeus/ebike3/packages/stations/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/stations/domain/port/in"
	"github.com/llamadeus/ebike3/packages/stations/infrastructure/micro"
)

func MakeStationsHandler(stationService in.StationService) micro.HTTPHandler {
	return micro.MakeHandler(func(ctx micro.Context[any, any]) ([]*dto.StationDTO, error) {
		stations, err := stationService.GetStationViews()
		if err != nil {
			return nil, err
		}

		dtos := make([]*dto.StationDTO, len(stations))
		for i, station := range stations {
			dtos[i] = dto.StationViewToDTO(station)
		}

		return dtos, nil
	})
}
