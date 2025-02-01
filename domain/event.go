package domain

import "github.com/google/uuid"

type Event interface {
	Id() uuid.UUID
	Name() string
}
