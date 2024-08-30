package slice

import (
	"context"

	"github.com/renatocosta55sp/modeling/domain"
)

type EventHandleable interface {
	Handle(ctx context.Context, event domain.Event) error
}
