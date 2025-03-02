package in

import (
	"github.com/llamadeus/ebike3/packages/vehicles/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/port/in"
	"github.com/llamadeus/ebike3/packages/vehicles/infrastructure/micro"
)

func MakeDeleteVehicleHandler(vehicleService in.VehicleService) micro.HTTPHandler {
	type params struct {
		ID string `param:"id"`
	}

	return micro.MakeHandler(func(ctx micro.Context[params, any]) (*dto.VehicleDTO, error) {
		id, err := dto.IDFromDTO(ctx.Params().ID)
		if err != nil {
			return nil, micro.NewBadRequestError("invalid vehicle id")
		}

		vehicle, err := vehicleService.DeleteVehicle(id)
		if vehicle == nil {
			return nil, err
		}

		return dto.VehicleToDTO(vehicle), nil
	})
}
