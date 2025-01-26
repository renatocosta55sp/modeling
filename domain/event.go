package domain

import (
	"encoding/json"
	"time"
)

// Event represents a generic event structure
type Event struct {
	Type      string
	Timestamp time.Time
	Data      any
	Metadata  any
}

func (e *Event) Serialize() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Event) Unserialize(data []byte) error {
	return json.Unmarshal(data, e)
}
