package bus2

import "fmt"

// EventDispatcher is responsible for dispatching events to their respective handlers.
type EventDispatcher struct {
	handlers map[string]func(Event) error
}

// NewEventDispatcher creates a new EventDispatcher.
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string]func(Event) error),
	}
}

// RegisterHandler registers an event handler for a specific event type.
func RegisterHandler[T Event](dispatcher *EventDispatcher, handler EventHandler[T]) {
	var event T // Create a zero-value instance of T to get the event name
	eventName := event.EventName()
	dispatcher.handlers[eventName] = func(event Event) error {
		return handler.Handle(event.(T))
	}
}

// Dispatch dispatches an event to its registered handler.
func (d *EventDispatcher) Dispatch(event Event) error {
	if handler, ok := d.handlers[event.EventName()]; ok {
		return handler(event)
	}
	return fmt.Errorf("no handler registered for event: %s", event.EventName())
}
