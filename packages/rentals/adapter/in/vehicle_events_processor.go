package in

import (
	"github.com/llamadeus/ebike3/packages/rentals/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/rentals/domain/events"
	"github.com/llamadeus/ebike3/packages/rentals/domain/port/in"
	"github.com/llamadeus/ebike3/packages/rentals/infrastructure/micro"
	"log/slog"
)

func MakeVehicleEventsProcessor(vehicleService in.VehicleService) *micro.EventsProcessor {
	return micro.NewEventsProcessor(micro.HandlersMap{
		events.VehiclesVehicleCreatedEventType: micro.NewEventHandler(func(payload events.VehicleCreatedEvent) error {
			slog.Info(
				"vehicle created",
				"id", payload.ID,
				"type", payload.Type,
				"positionX", payload.PositionX,
				"positionY", payload.PositionY,
				"battery", payload.Battery,
				"available", payload.Available,
			)

			id, err := dto.IDFromDTO(payload.ID)
			if err != nil {
				return err
			}

			type_, err := dto.TypeFromDTO(payload.Type)
			if err != nil {
				return err
			}

			return vehicleService.CreateVehicleView(id, type_)
		}),
		events.VehiclesVehicleDeletedEventType: micro.NewEventHandler(func(payload events.VehicleDeletedEvent) error {
			slog.Info(
				"vehicle deleted",
				"id", payload.ID,
			)

			id, err := dto.IDFromDTO(payload.ID)
			if err != nil {
				return err
			}

			return vehicleService.DeleteVehicleView(id)
		}),
	})
}
