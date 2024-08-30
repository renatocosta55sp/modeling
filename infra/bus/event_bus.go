package bus

import (
	"sync"

	"github.org/renatocosta55sp/modeling/domain"
)

// EventBus represents the event bus that handles event subscription and dispatching
type EventBus struct {
	mu           sync.RWMutex
	Subscribers  map[string][]chan<- domain.Event
	EventsRaised map[string]string
	Err          error
}

// NewEventBus creates a new instance of the event bus
func NewEventBus() *EventBus {
	return &EventBus{
		Subscribers:  make(map[string][]chan<- domain.Event),
		EventsRaised: make(map[string]string),
	}
}

// Subscribe adds a new subscriber for a given event type
func (eb *EventBus) Subscribe(eventType string, subscriber chan<- domain.Event) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.Subscribers[eventType] = append(eb.Subscribers[eventType], subscriber)
}

func (eb *EventBus) Publish(event domain.Event) {

	eb.mu.RLock()
	defer eb.mu.RUnlock()

	subscribers := eb.Subscribers[event.Type]

	for _, subscriber := range subscribers {
		subscriber <- event
		eb.EventsRaised[event.Type] = event.Type
	}

}

func (eb *EventBus) RaisedEvents() map[string]string {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	return eb.EventsRaised
}
