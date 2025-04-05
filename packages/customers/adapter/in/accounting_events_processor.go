package in

import (
	"github.com/llamadeus/ebike3/packages/customers/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/customers/domain/events"
	"github.com/llamadeus/ebike3/packages/customers/domain/port/in"
	"github.com/llamadeus/ebike3/packages/customers/infrastructure/micro"
	"log/slog"
)

func MakeAccountingEventsProcessor(customerService in.CustomerService) *micro.EventsProcessor {
	return micro.NewEventsProcessor(micro.HandlersMap{
		events.AccountingPaymentUpdatedEventType: micro.NewEventHandler(func(payload events.PaymentUpdatedEvent) error {
			slog.Info(
				"payment updated",
				"id", payload.ID,
				"customerId", payload.CustomerID,
				"amount", payload.Amount,
				"status", payload.Status,
			)

			if payload.Status != "CONFIRMED" {
				return nil
			}

			customerID, err := dto.IDFromDTO(payload.CustomerID)
			if err != nil {
				return err
			}

			return customerService.UpdateCustomerViewCreditBalance(customerID, payload.Amount)
		}),
		events.AccountingPaymentDeletedEventType: micro.NewEventHandler(func(payload events.PaymentDeletedEvent) error {
			slog.Info(
				"payment deleted",
				"id", payload.ID,
				"customerId", payload.CustomerID,
				"amount", payload.Amount,
				"status", payload.Status,
			)

			if payload.Status != "CONFIRMED" {
				return nil
			}

			customerID, err := dto.IDFromDTO(payload.CustomerID)
			if err != nil {
				return err
			}

			return customerService.UpdateCustomerViewCreditBalance(customerID, -payload.Amount)
		}),
		events.AccountingExpenseCreatedEventType: micro.NewEventHandler(func(payload events.ExpenseCreatedEvent) error {
			slog.Info(
				"expense created",
				"id", payload.ID,
				"customerId", payload.CustomerID,
				"rentalId", payload.RentalID,
				"amount", payload.Amount,
			)

			customerID, err := dto.IDFromDTO(payload.CustomerID)
			if err != nil {
				return err
			}

			return customerService.UpdateCustomerViewCreditBalance(customerID, -payload.Amount)
		}),
	})
}
