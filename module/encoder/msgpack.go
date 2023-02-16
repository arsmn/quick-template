package encoder

import "github.com/vmihailenco/msgpack/v5"

type MsgPackEncoder[T any] struct{}

var name MsgPackEncoder[any]

func (e MsgPackEncoder[T]) Encode(v T) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (e MsgPackEncoder[T]) Decode(data []byte) (T, error) {
	var v T
	err := msgpack.Unmarshal(data, &v)
	return v, err
}
