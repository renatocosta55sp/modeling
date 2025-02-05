package bus

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/renatocosta55sp/modeling/domain"
)

// EventRegistry is a component responsible for managing event types and creating event instances.
type EventRegistry struct {
	typeMap map[string]reflect.Type
}

func NewEventRegistry() *EventRegistry {
	return &EventRegistry{
		typeMap: make(map[string]reflect.Type),
	}
}

func (r *EventRegistry) RegisterEvents(events map[string]reflect.Type) {
	for eventName, eventType := range events {
		r.typeMap[eventName] = eventType
	}
}

// CreateEvent creates a new instance of the event based on the event type and unmarshals the payload.
func (r *EventRegistry) CreateEvent(eventType string, payload []byte) (*domain.Event, error) {
	eventTypeInstance, ok := r.typeMap[eventType]
	if !ok {
		return nil, fmt.Errorf("unknown event type: %s", eventType)
	}

	// Create a new instance of the event type
	event := reflect.New(eventTypeInstance).Interface()

	// Unmarshal the jsonb payload into the event instance
	if err := json.Unmarshal(payload, event); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event payload: %w", err)
	}

	// Cast the event to domain.Event
	if domainEvent, ok := event.(*domain.Event); ok {
		return domainEvent, nil
	}

	return nil, fmt.Errorf("event does not implement domain.Event interface")
}
