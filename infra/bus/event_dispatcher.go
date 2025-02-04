package bus

import (
	"fmt"
	"reflect"

	"github.com/renatocosta55sp/modeling/domain"
)

// EventDispatcher is responsible for dispatching events to their respective handlers.
type EventDispatcher struct {
	handlers map[string]func(domain.Event) error
}

// NewEventDispatcher creates a new EventDispatcher.
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string]func(domain.Event) error),
	}
}

// RegisterHandler registers an event handler for a specific event type.
func RegisterHandler[T domain.Event](dispatcher *EventDispatcher, handler EventHandler[T]) {

	eventType := reflect.TypeOf((*T)(nil)).Elem() // Get the concrete type of T
	eventName := eventType.Name()

	dispatcher.handlers[eventName] = func(event domain.Event) error {
		return handler.Handle(event.(T))
	}
}

// Dispatch dispatches an event to its registered handler.
func (d *EventDispatcher) Dispatch(event domain.Event) error {
	if handler, ok := d.handlers[event.GetName()]; ok {
		return handler(event)
	}
	return fmt.Errorf("no handler registered for event: %s", event.GetName())
}

func (d *EventDispatcher) DispatchUncommittedEvents(events []domain.Event) error {

	for _, evt := range events {
		if err := d.Dispatch(evt); err != nil {
			return err
		}
	}

	return nil
}
