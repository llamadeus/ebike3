package in

import (
	"github.com/llamadeus/ebike3/packages/accounting/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/accounting/domain/port/in"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/micro"
)

func MakeCreateExpenseHandler(accountingService in.AccountingService) micro.HTTPHandler {
	type input struct {
		CustomerID string `json:"customerId" validate:"required"`
		RentalID   string `json:"rentalId" validate:"required"`
		Amount     int32  `json:"amount" validate:"required"`
	}

	return micro.MakeHandler(func(ctx micro.Context[any, input]) (*dto.ExpenseDTO, error) {
		customerID, err := dto.IDFromDTO(ctx.Input().CustomerID)
		if err != nil {
			return nil, err
		}

		rentalID, err := dto.IDFromDTO(ctx.Input().RentalID)
		if err != nil {
			return nil, err
		}

		expense, err := accountingService.CreateExpense(customerID, rentalID, ctx.Input().Amount)
		if expense == nil {
			return nil, err
		}

		return dto.ExpenseToDTO(expense), nil
	})
}
