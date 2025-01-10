package domain

import (
	"github.com/google/uuid"
)

type Aggregate struct {
	AggregateID       uuid.UUID
	Version           int8
	UncommittedEvents []Event `json:"-"`
}
