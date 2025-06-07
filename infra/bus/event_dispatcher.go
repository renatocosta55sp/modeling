package bus

import (
	"fmt"

	"github.com/renatocosta55sp/modeling/domain"
)

type EventDispatcher struct {
	handlers map[string][]func(domain.Event) error
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]func(domain.Event) error),
	}
}

func RegisterHandler[T domain.Event](dispatcher *EventDispatcher, handler EventHandler[T]) {
	var event T
	eventName := event.GetName()
	dispatcher.handlers[eventName] = append(dispatcher.handlers[eventName], func(event domain.Event) error {
		return handler.Handle(event.(T))
	})
}

func (d *EventDispatcher) Dispatch(event domain.Event) error {
	if handlers, ok := d.handlers[event.GetName()]; ok {
		for _, handler := range handlers {
			if err := handler(event); err != nil {
				return err
			}
		}
		return nil
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
