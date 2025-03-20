package in

import (
	"github.com/llamadeus/ebike3/packages/rentals/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/rentals/domain/events"
	"github.com/llamadeus/ebike3/packages/rentals/domain/port/in"
	"github.com/llamadeus/ebike3/packages/rentals/infrastructure/micro"
	"log/slog"
)

func MakeAccountingEventsProcessor(rentalService in.RentalService) *micro.EventsProcessor {
	return micro.NewEventsProcessor(micro.HandlersMap{
		events.AccountingExpenseCreatedEventType: micro.NewEventHandler(func(payload events.ExpenseCreatedEvent) error {
			slog.Info(
				"expense created",
				"id", payload.ID,
				"customerId", payload.CustomerID,
				"rentalId", payload.RentalID,
				"amount", payload.Amount,
			)

			rentalID, err := dto.IDFromDTO(payload.RentalID)
			if err != nil {
				return err
			}

			return rentalService.AddExpenseToRental(rentalID, payload.Amount)
		}),
	})
}
