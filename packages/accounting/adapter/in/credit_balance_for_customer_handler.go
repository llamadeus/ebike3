package in

import (
	"github.com/llamadeus/ebike3/packages/accounting/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/accounting/domain/port/in"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/micro"
)

func MakeGetCreditBalanceForCustomerHandler(accountingService in.AccountingService) micro.HTTPHandler {
	type params struct {
		ID string `param:"id"`
	}

	return micro.MakeHandler(func(ctx micro.Context[params, any]) (*dto.CreditBalanceDTO, error) {
		id, err := dto.IDFromDTO(ctx.Params().ID)
		if err != nil {
			return nil, err
		}

		creditBalance, err := accountingService.GetCreditBalanceForCustomer(id)
		if err != nil {
			return nil, err
		}

		return dto.CreditBalanceToDTO(id, creditBalance), nil
	})
}
