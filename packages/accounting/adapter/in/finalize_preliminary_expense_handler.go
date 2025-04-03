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

	type input struct {
		RentalID string `json:"rentalId" validate:"required"`
	}

	return micro.MakeHandler(func(ctx micro.Context[params, input]) (*dto.ExpenseDTO, error) {
		id, err := dto.IDFromDTO(ctx.Params().ID)
		if err != nil {
			return nil, err
		}

		rentalID, err := dto.IDFromDTO(ctx.Input().RentalID)
		if err != nil {
			return nil, err
		}

		expense, err := accountingService.FinalizePreliminaryExpense(id, rentalID)
		if expense == nil {
			return nil, err
		}

		return dto.ExpenseToDTO(expense), nil
	})
}
