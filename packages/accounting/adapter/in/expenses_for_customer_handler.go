package in

import (
	"github.com/llamadeus/ebike3/packages/accounting/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/accounting/domain/port/in"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/micro"
)

func MakeGetExpensesForCustomerHandler(accountingService in.AccountingService) micro.HTTPHandler {
	type params struct {
		ID string `param:"id"`
	}

	return micro.MakeHandler(func(ctx micro.Context[params, any]) ([]*dto.ExpenseDTO, error) {
		id, err := dto.IDFromDTO(ctx.Params().ID)
		if err != nil {
			return nil, err
		}

		expenses, err := accountingService.GetExpensesForCustomer(id)
		if err != nil {
			return nil, err
		}

		dtos := make([]*dto.ExpenseDTO, len(expenses))
		for i, expense := range expenses {
			dtos[i] = dto.ExpenseToDTO(expense)
		}

		return dtos, nil
	})
}
