package in

import (
	"github.com/llamadeus/ebike3/packages/rentals/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/rentals/domain/port/in"
	"github.com/llamadeus/ebike3/packages/rentals/infrastructure/micro"
)

func MakeStopRentalHandler(rentalService in.RentalService) micro.HTTPHandler {
	type params struct {
		ID string `param:"id"`
	}

	type input struct {
		CustomerID string `json:"customerId" validate:"required"`
	}

	return micro.MakeHandler(func(ctx micro.Context[params, input]) (*dto.RentalDTO, error) {
		id, err := dto.IDFromDTO(ctx.Params().ID)
		if err != nil {
			return nil, micro.NewBadRequestError("invalid rental id")
		}

		customerID, err := dto.IDFromDTO(ctx.Input().CustomerID)
		if err != nil {
			return nil, micro.NewBadRequestError("invalid customer id")
		}

		rental, err := rentalService.StopRental(id, customerID)
		if rental == nil {
			return nil, err
		}

		rentalView, err := rentalService.GetRentalView(id)
		if err != nil {
			return nil, err
		}

		return dto.RentalToDTO(rental, rentalView.Cost), nil
	})
}
