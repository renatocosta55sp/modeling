package bus

import "github.com/renatocosta55sp/modeling/domain"

type EventResult struct {
	Event domain.Event
	Err   error
}
