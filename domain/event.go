package domain

import "github.com/google/uuid"

type Event interface {
	GetId() uuid.UUID
	GetName() string
}
