package in

import (
	"github.com/llamadeus/ebike3/packages/accounting/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/accounting/domain/port/in"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/micro"
)

func MakeGetPaymentsHandler(accountingService in.AccountingService) micro.HTTPHandler {
	return micro.MakeHandler(func(ctx micro.Context[any, any]) ([]*dto.PaymentDTO, error) {
		payments, err := accountingService.GetAllPayments()
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
