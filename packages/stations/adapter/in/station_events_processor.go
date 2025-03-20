package in

import (
	"github.com/llamadeus/ebike3/packages/stations/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/stations/domain/events"
	"github.com/llamadeus/ebike3/packages/stations/domain/port/in"
	"github.com/llamadeus/ebike3/packages/stations/infrastructure/micro"
	"log/slog"
)

func MakeStationEventsProcessor(stationService in.StationService) *micro.EventsProcessor {
	return micro.NewEventsProcessor(micro.HandlersMap{
		events.StationsStationCreatedEventType: micro.NewEventHandler(func(payload events.StationCreatedEvent) error {
			slog.Info(
				"station created",
				"id", payload.ID,
				"name", payload.Name,
				"positionX", payload.PositionX,
				"positionY", payload.PositionY,
			)

			id, err := dto.IDFromDTO(payload.ID)
			if err != nil {
				return err
			}

			return stationService.CreateStationView(id, payload.Name, payload.PositionX, payload.PositionY)
		}),
		events.StationsStationUpdatedEventType: micro.NewEventHandler(func(payload events.StationUpdatedEvent) error {
			slog.Info(
				"station updated",
				"id", payload.ID,
				"name", payload.Name,
				"positionX", payload.PositionX,
				"positionY", payload.PositionY,
			)

			id, err := dto.IDFromDTO(payload.ID)
			if err != nil {
				return err
			}

			return stationService.UpdateStationView(id, payload.Name, payload.PositionX, payload.PositionY)
		}),
		events.StationsStationDeletedEventType: micro.NewEventHandler(func(payload events.StationDeletedEvent) error {
			slog.Info(
				"station deleted",
				"id", payload.ID,
			)

			id, err := dto.IDFromDTO(payload.ID)
			if err != nil {
				return err
			}

			return stationService.DeleteStationView(id)
		}),
	})
}
