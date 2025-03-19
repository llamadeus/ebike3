package in

import (
	"github.com/llamadeus/ebike3/packages/accounting/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/accounting/domain/port/in"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/micro"
)

func MakeGetPaymentsForCustomerHandler(accountingService in.AccountingService) micro.HTTPHandler {
	type params struct {
		ID string `param:"id"`
	}

	return micro.MakeHandler(func(ctx micro.Context[params, any]) ([]*dto.PaymentDTO, error) {
		id, err := dto.IDFromDTO(ctx.Params().ID)
		if err != nil {
			return nil, err
		}

		payments, err := accountingService.GetPaymentsForCustomer(id)
		if err != nil {
			return nil, err
		}

		dtos := make([]*dto.PaymentDTO, len(payments))
		for i, station := range payments {
			dtos[i] = dto.PaymentToDTO(station)
		}

		return dtos, nil
	})
}
