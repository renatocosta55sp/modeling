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
	WriteSnapshot(ctx context.Context, streamId string, newEvents []domain.Event, expectedVersion int) error
	ReadSnapshot(streamID string) (domain.Event, error)
}
