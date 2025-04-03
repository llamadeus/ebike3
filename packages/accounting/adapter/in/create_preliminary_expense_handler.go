package in

import (
	"github.com/llamadeus/ebike3/packages/accounting/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/accounting/domain/port/in"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/micro"
)

func MakeCreatePreliminaryExpenseHandler(accountingService in.AccountingService) micro.HTTPHandler {
	type input struct {
		InquiryID  string `json:"inquiryId" validate:"required"`
		CustomerID string `json:"customerId" validate:"required"`
		RentalID   string `json:"rentalId" validate:"required"`
		Amount     int32  `json:"amount" validate:"required"`
	}

	return micro.MakeHandler(func(ctx micro.Context[any, input]) (*dto.PreliminaryExpenseDTO, error) {
		inquiryID, err := dto.IDFromDTO(ctx.Input().InquiryID)
		if err != nil {
			return nil, err
		}

		customerID, err := dto.IDFromDTO(ctx.Input().CustomerID)
		if err != nil {
			return nil, err
		}

		rentalID, err := dto.IDFromDTO(ctx.Input().RentalID)
		if err != nil {
			return nil, err
		}

		preliminaryExpense, err := accountingService.CreatePreliminaryExpense(inquiryID, customerID, rentalID, ctx.Input().Amount)
		if preliminaryExpense == nil {
			return nil, err
		}

		return dto.PreliminaryExpenseToDTO(preliminaryExpense), nil
	})
}
