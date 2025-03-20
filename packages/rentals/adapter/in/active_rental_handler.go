package in

import (
	"github.com/llamadeus/ebike3/packages/rentals/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/rentals/domain/port/in"
	"github.com/llamadeus/ebike3/packages/rentals/infrastructure/micro"
)

func MakeGetActiveRentalHandler(rentalService in.RentalService) micro.HTTPHandler {
	type params struct {
		ID string `param:"id"`
	}

	return micro.MakeHandler(func(ctx micro.Context[params, any]) (*dto.RentalDTO, error) {
		id, err := dto.IDFromDTO(ctx.Params().ID)
		if err != nil {
			return nil, micro.NewBadRequestError("invalid customer id")
		}

		rental, err := rentalService.GetActiveRentalForCustomer(id)
		if rental == nil {
			return nil, err
		}

		return dto.RentalViewToDTO(rental), nil
	})
}
