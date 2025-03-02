package in

import (
	"github.com/llamadeus/ebike3/packages/vehicles/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/port/in"
	"github.com/llamadeus/ebike3/packages/vehicles/infrastructure/micro"
)

func MakeAvailableVehiclesHandler(vehicleService in.VehicleService) micro.HTTPHandler {
	return micro.MakeHandler(func(ctx micro.Context[any, any]) ([]*dto.VehicleDTO, error) {
		vehicles, err := vehicleService.GetAvailableVehicleViews()
		if err != nil {
			return nil, err
		}

		dtos := make([]*dto.VehicleDTO, len(vehicles))
		for i, vehicle := range vehicles {
			dtos[i] = dto.VehicleViewToDTO(vehicle)
		}

		return dtos, nil
	})
}
