package in

import (
	"github.com/llamadeus/ebike3/packages/rentals/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/rentals/domain/port/in"
	"github.com/llamadeus/ebike3/packages/rentals/infrastructure/micro"
)

func MakeStartRentalHandler(rentalService in.RentalService) micro.HTTPHandler {
	type input struct {
		CustomerID string `json:"customerId" validate:"required"`
		VehicleID  string `json:"vehicleId" validate:"required"`
	}

	return micro.MakeHandler(func(ctx micro.Context[any, input]) (*dto.RentalDTO, error) {
		customerID, err := dto.IDFromDTO(ctx.Input().CustomerID)
		if err != nil {
			return nil, micro.NewBadRequestError("invalid customer id")
		}

		vehicleID, err := dto.IDFromDTO(ctx.Input().VehicleID)
		if err != nil {
			return nil, micro.NewBadRequestError("invalid vehicle id")
		}

		rental, err := rentalService.StartRental(customerID, vehicleID)
		if rental == nil {
			return nil, err
		}

		return dto.RentalToDTO(rental), nil
	})
}
