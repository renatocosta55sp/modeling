package bus2

import "github.com/renatocosta55sp/modeling/domain"

type EventHandler[T domain.Event] interface {
	Handle(event T) error
}
