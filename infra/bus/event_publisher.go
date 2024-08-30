package bus

import "github.org/renatocosta55sp/modeling/domain"

type EventPublisher struct {
	eventBus *EventBus
}

func NewEventPublisher(eventBus *EventBus) *EventPublisher {
	return &EventPublisher{
		eventBus: eventBus,
	}
}

func (ep *EventPublisher) Publish(events []domain.Event) {

	for _, event := range events {
		ep.eventBus.Publish(event)
	}
}
