package infra

import "encoding/json"

func Serialize(data any) ([]byte, error) {
	return json.Marshal(data)
}

func Unserialize(dataIn any, data []byte) error {
	return json.Unmarshal(data, dataIn)
}
