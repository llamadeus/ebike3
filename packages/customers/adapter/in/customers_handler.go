package in

import (
	"github.com/llamadeus/ebike3/packages/customers/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/customers/domain/port/in"
	"github.com/llamadeus/ebike3/packages/customers/infrastructure/micro"
)

func MakeCustomersHandler(customerService in.CustomerService) micro.HTTPHandler {
	return micro.MakeHandler(func(ctx micro.Context[any, any]) ([]*dto.CustomerDTO, error) {
		customers, err := customerService.GetCustomerViews()
		if err != nil {
			return nil, err
		}

		dtos := make([]*dto.CustomerDTO, len(customers))
		for i, customer := range customers {
			dtos[i] = dto.CustomerViewToDTO(customer)
		}

		return dtos, nil
	})
}
