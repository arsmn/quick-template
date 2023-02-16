package hash

import (
	"context"
	"errors"
)

var (
	ErrInvalidHash           = errors.New("the encoded hash is not in the correct format")
	ErrMismatchedHashAndData = errors.New("hash and data do not match")
)

type Hasher interface {
	Hash(ctx context.Context, data []byte) ([]byte, error)
	Compare(ctx context.Context, data []byte, hash []byte) error
}
