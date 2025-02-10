package eventstore

import (
	"context"

	"github.com/renatocosta55sp/modeling/domain"
)

func ShouldTakeSnapshot(agg domain.Aggregate, frequency int) bool {
	if frequency <= 0 {
		// Invalid frequency, default to not taking snapshots
		return false
	}
	return agg.Version%frequency == 0
}

type SnapshotStore interface {
	WriteSnapshot(ctx context.Context, streamId string, event any, version int, eventType string, metaData any, eventIdentifier string) error
	ReadSnapshot(ctx context.Context, streamID string) (domain.Event, error)
}
