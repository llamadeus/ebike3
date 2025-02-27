package in

import (
	"github.com/llamadeus/ebike3/packages/auth/domain/events"
	"github.com/llamadeus/ebike3/packages/auth/infrastructure/micro"
	"log/slog"
)

var EventsProcessor = micro.NewEventsProcessor(micro.HandlersMap{
	events.AuthUserRegisteredEventType: micro.NewEventHandler(func(payload events.UserRegisteredEvent) error {
		slog.Info(
			"user registered",
			"id", payload.ID,
			"username", payload.Username,
			"role", payload.Role,
		)

		return nil
	}),
})
