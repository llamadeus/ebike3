package in

import (
	"github.com/llamadeus/ebike3/packages/customers/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/customers/domain/events"
	"github.com/llamadeus/ebike3/packages/customers/domain/port/in"
	"github.com/llamadeus/ebike3/packages/customers/infrastructure/micro"
	"log/slog"
)

func MakeCustomerEventsProcessor(customerService in.CustomerService) *micro.EventsProcessor {
	return micro.NewEventsProcessor(micro.HandlersMap{
		events.CustomersCustomerPositionUpdatedEventType: micro.NewEventHandler(func(payload events.CustomerPositionUpdatedEvent) error {
			slog.Info(
				"customer position updated",
				"id", payload.ID,
				"positionX", payload.PositionX,
				"positionY", payload.PositionY,
			)

			id, err := dto.IDFromDTO(payload.ID)
			if err != nil {
				return err
			}

			return customerService.UpdateCustomerViewPosition(id, payload.PositionX, payload.PositionY)
		}),
	})
}
