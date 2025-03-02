package in

import (
	"github.com/llamadeus/ebike3/packages/customers/domain/port/in"
	"github.com/llamadeus/ebike3/packages/customers/infrastructure/micro"
)

func MakePaymentEventsProcessor(customerService in.CustomerService) *micro.EventsProcessor {
	return micro.NewEventsProcessor(micro.HandlersMap{
		// TODO: Handle payment confirmed event to increase credit balance
	})
}
