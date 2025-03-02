package in

import (
	"github.com/llamadeus/ebike3/packages/customers/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/customers/domain/events"
	"github.com/llamadeus/ebike3/packages/customers/domain/port/in"
	"github.com/llamadeus/ebike3/packages/customers/infrastructure/micro"
	"log/slog"
)

func MakeAuthEventsProcessor(customerService in.CustomerService) *micro.EventsProcessor {
	return micro.NewEventsProcessor(micro.HandlersMap{
		events.AuthUserRegisteredEventType: micro.NewEventHandler(func(payload events.UserRegisteredEvent) error {
			slog.Info(
				"user registered",
				"id", payload.ID,
				"username", payload.Username,
				"role", payload.Role,
			)

			if payload.Role != "CUSTOMER" {
				return nil
			}

			id, err := dto.IDFromDTO(payload.ID)
			if err != nil {
				return err
			}

			return customerService.CreateCustomerView(id, payload.Username)
		}),
		events.AuthUserLoggedInEventType: micro.NewEventHandler(func(payload events.UserLoggedInEvent) error {
			slog.Info(
				"user logged in",
				"id", payload.ID,
				"username", payload.Username,
				"timestamp", payload.Timestamp,
			)

			id, err := dto.IDFromDTO(payload.ID)
			if err != nil {
				return err
			}

			return customerService.UpdateCustomerViewLastLogin(id, payload.Timestamp)
		}),
	})
}
