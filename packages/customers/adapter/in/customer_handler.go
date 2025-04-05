package in

import (
	"github.com/llamadeus/ebike3/packages/customers/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/customers/domain/port/in"
	"github.com/llamadeus/ebike3/packages/customers/infrastructure/micro"
)

func MakeCustomerHandler(customerService in.CustomerService) micro.HTTPHandler {
	type params struct {
		ID string `param:"id"`
	}

	return micro.MakeHandler(func(ctx micro.Context[params, any]) (*dto.CustomerDTO, error) {
		id, err := dto.IDFromDTO(ctx.Params().ID)
		if err != nil {
			return nil, micro.NewBadRequestError("invalid customer id")
		}

		customer, err := customerService.GetCustomerViewByID(id)
		if err != nil {
			return nil, err
		}

		return dto.CustomerViewToDTO(customer), nil
	})
}
