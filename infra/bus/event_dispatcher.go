package bus

import (
	"github.com/gookit/event"
	"github.com/renatocosta55sp/modeling/domain"
)

func DispatchUncommittedEvents(events []domain.Event, agg *domain.Aggregate) error {

	for _, evt := range events {
		eventHandlerResult := event.MustFire(evt.Type, event.M{"event": evt.Data, "agg": agg})

		if errValue, ok := eventHandlerResult.Get("error").(error); ok {
			return errValue
		}
	}

	return nil
}
