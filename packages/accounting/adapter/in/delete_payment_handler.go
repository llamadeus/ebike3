package in

import (
	"github.com/llamadeus/ebike3/packages/accounting/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/accounting/domain/port/in"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/micro"
)

func MakeDeletePaymentHandler(accountingService in.AccountingService) micro.HTTPHandler {
	type params struct {
		ID string `param:"id"`
	}

	type input struct {
		CustomerID string `json:"customerId" validate:"required"`
	}

	return micro.MakeHandler(func(ctx micro.Context[params, input]) (*dto.PaymentDTO, error) {
		id, err := dto.IDFromDTO(ctx.Params().ID)
		if err != nil {
			return nil, err
		}

		customerID, err := dto.IDFromDTO(ctx.Input().CustomerID)
		if err != nil {
			return nil, err
		}

		payment, err := accountingService.DeleteCustomerPayment(id, customerID)
		if payment == nil {
			return nil, err
		}

		return dto.PaymentToDTO(payment), nil
	})
}
