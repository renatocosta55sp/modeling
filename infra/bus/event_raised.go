package bus

import "github.org/renatocosta55sp/modeling/domain"

type EventResult struct {
	Event domain.Event
	Err   error
}
