package in

import (
	"github.com/llamadeus/ebike3/packages/accounting/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/accounting/domain/port/in"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/micro"
)

func MakeFinalizePreliminaryExpenseHandler(accountingService in.AccountingService) micro.HTTPHandler {
	type params struct {
		ID string `param:"id"`
	}

	return micro.MakeHandler(func(ctx micro.Context[params, any]) (*dto.ExpenseDTO, error) {
		id, err := dto.IDFromDTO(ctx.Params().ID)
		if err != nil {
			return nil, err
		}

		expense, err := accountingService.FinalizePreliminaryExpense(id)
		if expense == nil {
			return nil, err
		}

		return dto.ExpenseToDTO(expense), nil
	})
}
