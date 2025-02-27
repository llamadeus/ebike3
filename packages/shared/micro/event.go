package micro

import (
	"encoding/json"
	"time"
)

type (
	Event[T any] struct {
		Type      string    `json:"type"`
		Payload   T         `json:"payload"`
		Timestamp time.Time `json:"timestamp"`
	}

	EventHandler func(event Event[json.RawMessage]) error

	HandlersMap map[string]EventHandler

	EventsProcessor struct {
		handlers HandlersMap
	}
)

func NewEvent[T any](eventType string, payload T) Event[T] {
	return Event[T]{
		Type:      eventType,
		Payload:   payload,
		Timestamp: time.Now(),
	}
}

func NewEventHandler[T any](handler func(payload T) error) EventHandler {
	return func(event Event[json.RawMessage]) error {
		var payload T
		err := json.Unmarshal(event.Payload, &payload)
		if err != nil {
			return err
		}

		return handler(payload)
	}
}

func NewEventsProcessor(handlers HandlersMap) *EventsProcessor {
	return &EventsProcessor{
		handlers: handlers,
	}
}

func (p *EventsProcessor) Handle(event Event[json.RawMessage]) error {
	handler, ok := p.handlers[event.Type]
	if !ok {
		return nil
	}

	return handler(event)
}
