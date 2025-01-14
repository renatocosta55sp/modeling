package eventsourcing

import "github.com/renatocosta55sp/modeling/domain"

func ShouldTakeSnapshot(agg domain.Aggregate, frequency int) bool {
	if frequency <= 0 {
		// Invalid frequency, default to not taking snapshots
		return false
	}
	return agg.Version%frequency == 0
}
