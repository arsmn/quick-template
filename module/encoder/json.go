package encoder

import "encoding/json"

type JSONEncoder[T any] struct{}

func (e JSONEncoder[T]) Encode(v T) ([]byte, error) {
	return json.Marshal(v)
}

func (e JSONEncoder[T]) Decode(data []byte) (T, error) {
	var v T
	err := json.Unmarshal(data, &v)
	return v, err
}
