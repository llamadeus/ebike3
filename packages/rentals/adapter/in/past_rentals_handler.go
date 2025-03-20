package in

import (
	"github.com/llamadeus/ebike3/packages/rentals/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/rentals/domain/port/in"
	"github.com/llamadeus/ebike3/packages/rentals/infrastructure/micro"
)

func MakeGetPastRentalsHandler(rentalService in.RentalService) micro.HTTPHandler {
	type params struct {
		ID string `param:"id"`
	}

	return micro.MakeHandler(func(ctx micro.Context[params, any]) ([]*dto.RentalDTO, error) {
		id, err := dto.IDFromDTO(ctx.Params().ID)
		if err != nil {
			return nil, micro.NewBadRequestError("invalid customer id")
		}

		rentals, err := rentalService.GetPastRentalsForCustomer(id)
		if err != nil {
			return nil, err
		}

		dtos := make([]*dto.RentalDTO, len(rentals))
		for i, rental := range rentals {
			dtos[i] = dto.RentalViewToDTO(rental)
		}

		return dtos, nil
	})
}
