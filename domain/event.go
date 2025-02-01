package domain

import "github.com/google/uuid"

type Event interface {
	EventId() uuid.UUID
	EventName() string
}
