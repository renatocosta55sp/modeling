package slice

import (
	"context"

	"github.org/renatocosta55sp/modeling/domain"
)

type EventHandleable interface {
	Handle(ctx context.Context, event domain.Event) error
}
