package domain

import (
	"github.com/google/uuid"
)

type Aggregate struct {
	AggregateID       uuid.UUID
	Version           int
	UncommittedEvents []*Event `json:"-"`
}
