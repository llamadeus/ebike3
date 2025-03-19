package in

import (
	"github.com/llamadeus/ebike3/packages/accounting/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/accounting/domain/port/in"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/micro"
)

func MakeUpdatePaymentHandler(accountingService in.AccountingService) micro.HTTPHandler {
	type params struct {
		ID string `param:"id"`
	}

	type input struct {
		Status string `json:"status" validate:"required,oneof=CONFIRMED REJECTED"`
	}

	return micro.MakeHandler(func(ctx micro.Context[params, input]) (*dto.PaymentDTO, error) {
		id, err := dto.IDFromDTO(ctx.Params().ID)
		if err != nil {
			return nil, err
		}

		status, err := dto.StatusFromDTO(ctx.Input().Status)
		if err != nil {
			return nil, err
		}

		payment, err := accountingService.UpdatePayment(id, status)
		if payment == nil {
			return nil, err
		}

		return dto.PaymentToDTO(payment), nil
	})
}
