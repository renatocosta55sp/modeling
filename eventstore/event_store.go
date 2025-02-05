package eventstore

import (
	"context"
	"errors"

	"github.com/renatocosta55sp/modeling/domain"
)

var ErrConcurrencyConflict = errors.New("concurrency conflict: event version mismatch")

type EventStore interface {
	AppendToStream(ctx context.Context, streamId string, newEvents []domain.Event, expectedVersion int) error
	ReadStream(ctx context.Context, streamID string) ([]*domain.Event, error)
	ReadAllStream(ctx context.Context) ([]*domain.Event, error)
}
