package domain

type Event interface {
	EventName() string
}

/*
func (ev Event) Serialize(data any) ([]byte, error) {
	return json.Marshal(data)
}

func (*Event) Unserialize(dataIn any, data []byte) error {
	return json.Unmarshal(data, dataIn)
}*/
