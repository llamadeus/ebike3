package in

import (
	"github.com/llamadeus/ebike3/packages/accounting/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/accounting/domain/port/in"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/micro"
)

func MakeCreatePaymentHandler(accountingService in.AccountingService) micro.HTTPHandler {
	type input struct {
		CustomerID string `json:"customerId" validate:"required"`
		Amount     int    `json:"amount" validate:"required"`
	}

	return micro.MakeHandler(func(ctx micro.Context[any, input]) (*dto.PaymentDTO, error) {
		customerID, err := dto.IDFromDTO(ctx.Input().CustomerID)
		if err != nil {
			return nil, err
		}

		payment, err := accountingService.CreatePayment(customerID, ctx.Input().Amount)
		if payment == nil {
			return nil, err
		}

		return dto.PaymentToDTO(payment), nil
	})
}
