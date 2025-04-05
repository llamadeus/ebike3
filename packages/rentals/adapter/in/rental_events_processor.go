package in

import (
	"github.com/llamadeus/ebike3/packages/rentals/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/rentals/domain/events"
	"github.com/llamadeus/ebike3/packages/rentals/domain/port/in"
	"github.com/llamadeus/ebike3/packages/rentals/infrastructure/micro"
	"log/slog"
)

func MakeRentalEventsProcessor(rentalService in.RentalService) *micro.EventsProcessor {
	return micro.NewEventsProcessor(micro.HandlersMap{
		events.RentalsRentalStoppedEventType: micro.NewEventHandler(func(payload events.RentalStoppedEvent) error {
			slog.Info(
				"rental stopped",
				"id", payload.ID,
				"customerId", payload.CustomerID,
				"vehicleId", payload.VehicleID,
				"start", payload.Start,
				"end", payload.End,
			)

			id, err := dto.IDFromDTO(payload.ID)
			if err != nil {
				return err
			}

			return rentalService.UpdateRentalView(id, payload.End)
		}),
	})
}
