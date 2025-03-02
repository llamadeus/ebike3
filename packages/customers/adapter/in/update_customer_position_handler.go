package in

import (
	"github.com/llamadeus/ebike3/packages/customers/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/customers/domain/port/in"
	"github.com/llamadeus/ebike3/packages/customers/infrastructure/micro"
)

func MakeUpdateCustomerPositionHandler(customerService in.CustomerService) micro.HTTPHandler {
	type params struct {
		ID string `param:"id"`
	}

	type input struct {
		PositionX float64 `json:"positionX" validate:"required"`
		PositionY float64 `json:"positionY" validate:"required"`
	}

	return micro.MakeHandler(func(ctx micro.Context[params, input]) (any, error) {
		id, err := dto.IDFromDTO(ctx.Params().ID)
		if err != nil {
			return nil, micro.NewBadRequestError("invalid customer id")
		}

		return nil, customerService.UpdateCustomerPosition(id, ctx.Input().PositionX, ctx.Input().PositionY)
	})
}
