package encoder

type NopEncoder[T any] struct{}

func (e NopEncoder[T]) Encode(v T) ([]byte, error) {
	return nil, nil
}

func (e NopEncoder[T]) Decode(data []byte) (T, error) {
	var v T
	return v, nil
}
