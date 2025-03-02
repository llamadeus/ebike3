package in

import (
	"github.com/llamadeus/ebike3/packages/customers/domain/port/in"
	"github.com/llamadeus/ebike3/packages/customers/infrastructure/micro"
)

func MakeRentalEventsProcessor(customerService in.CustomerService) *micro.EventsProcessor {
	return micro.NewEventsProcessor(micro.HandlersMap{
		// TODO: Handle rental started event to update active rental

		// TODO: Handle rental ended event to update active rental

		// TODO: Handle rental tick event to decrease credit balance
	})
}
