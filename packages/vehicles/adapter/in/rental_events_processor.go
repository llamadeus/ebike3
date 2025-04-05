package in

import (
	"github.com/llamadeus/ebike3/packages/vehicles/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/events"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/port/in"
	"github.com/llamadeus/ebike3/packages/vehicles/infrastructure/micro"
	"log/slog"
)

func MakeRentalEventsProcessor(vehicleService in.VehicleService) *micro.EventsProcessor {
	return micro.NewEventsProcessor(micro.HandlersMap{
		events.RentalsRentalStartedEventType: micro.NewEventHandler(func(payload events.RentalStartedEvent) error {
			slog.Info(
				"rental started",
				"id", payload.ID,
				"customerId", payload.CustomerID,
				"vehicleId", payload.VehicleID,
				"vehicleType", payload.VehicleType,
				"start", payload.Start,
				"end", payload.End,
			)

			rentalID, err := dto.IDFromDTO(payload.ID)
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

			return vehicleService.UpdateVehicleViewActiveRental(rentalID, customerID, vehicleID, payload.VehicleType, payload.Start, 0)
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

			rentalID, err := dto.IDFromDTO(payload.ID)
			if err != nil {
				return err
			}

			customerID, err := dto.IDFromDTO(payload.CustomerID)
			if err != nil {
				return err
			}

			return vehicleService.ResetVehicleViewActiveRental(customerID, rentalID)
		}),
		events.RentalsCostUpdatedType: micro.NewEventHandler(func(payload events.CostUpdatedEvent) error {
			slog.Info(
				"cost updated",
				"id", payload.ID,
				"customerId", payload.CustomerID,
				"vehicleId", payload.VehicleID,
				"vehicleType", payload.VehicleType,
				"start", payload.Start,
				"cost", payload.Cost,
			)

			rentalID, err := dto.IDFromDTO(payload.ID)
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

			return vehicleService.UpdateVehicleViewActiveRental(rentalID, customerID, vehicleID, payload.VehicleType, payload.Start, payload.Cost)
		}),
	})
}
