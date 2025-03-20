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
		events.RentalsRentalStartedEventType: micro.NewEventHandler(func(payload events.RentalStartedEvent) error {
			slog.Info(
				"rental started",
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

			customerID, err := dto.IDFromDTO(payload.CustomerID)
			if err != nil {
				return err
			}

			vehicleID, err := dto.IDFromDTO(payload.VehicleID)
			if err != nil {
				return err
			}

			vehicleType, err := dto.TypeFromDTO(payload.VehicleType)
			if err != nil {
				return err
			}

			return rentalService.CreateRentalView(id, customerID, vehicleID, vehicleType, payload.Start)
		}),
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
