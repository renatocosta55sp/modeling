package bus

import (
	"github.com/gookit/event"
	"github.com/renatocosta55sp/modeling/domain"
)

func DispatchUncommittedEvents(events []domain.Event) error {

	for _, evt := range events {
		eventHandlerResult := event.MustFire(evt.Type, event.M{"data": evt.Data})

		if errValue, ok := eventHandlerResult.Get("error").(error); ok {
			return errValue
		}
	}

	return nil
}
