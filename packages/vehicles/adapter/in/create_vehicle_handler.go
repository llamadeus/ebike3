package in

import (
	"github.com/llamadeus/ebike3/packages/vehicles/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/model"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/port/in"
	"github.com/llamadeus/ebike3/packages/vehicles/infrastructure/micro"
)

func MakeCreateVehicleHandler(vehicleService in.VehicleService) micro.HTTPHandler {
	type input struct {
		Type      model.VehicleType `json:"type" validate:"required,oneof=BIKE EBIKE ABIKE"`
		PositionX float64           `json:"positionX"`
		PositionY float64           `json:"positionY"`
	}

	return micro.MakeHandler(func(ctx micro.Context[any, input]) (*dto.VehicleDTO, error) {
		vehicle, err := vehicleService.CreateVehicle(ctx.Input().Type, ctx.Input().PositionX, ctx.Input().PositionY, 1)
		if vehicle == nil {
			return nil, err
		}

		return dto.VehicleToDTO(vehicle), nil
	})
}
