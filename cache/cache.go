package cache

import (
	"context"
)

type Fallback[T any] func() (T, error)

type Cacher[T any] interface {
	Get(ctx context.Context, key any) (T, error)
	GetFallback(ctx context.Context, key any, fallback Fallback[T]) (T, error)
	Set(ctx context.Context, key any, value T) error
	Del(ctx context.Context, key any) error
}
