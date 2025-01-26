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

func (*Event) Serialize(data any) ([]byte, error) {
	return json.Marshal(data)
}

func (*Event) Unserialize(dataIn any, data []byte) error {
	return json.Unmarshal(data, dataIn)
}
