package redis

import (
	"QUICK-Template/cache"
	"QUICK-Template/module/encoder"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

type Cacher[T any] struct {
	prefix  string
	ttl     time.Duration
	client  *redis.Client
	encoder encoder.Encoder[T]
}

func NewCacher[T any](
	cli *redis.Client,
	prefix string,
	ttl time.Duration,
	format encoder.EncoderFormat,
) Cacher[T] {
	return Cacher[T]{
		client:  cli,
		ttl:     ttl,
		prefix:  prefix,
		encoder: encoder.GetEncoder[T](format),
	}
}

func (c Cacher[T]) key(key any) string {
	return fmt.Sprintf("%s_%v", c.prefix, key)
}

func (c Cacher[T]) Set(ctx context.Context, key any, item T) error {
	b, err := c.encoder.Encode(item)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, c.key(key), b, c.ttl).Err()
}

func (c Cacher[T]) Get(ctx context.Context, key any) (T, error) {
	var item T

	b, err := c.client.Get(ctx, c.key(key)).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return item, cache.ErrCacheMiss
		}
		return item, err
	}

	item, err = c.encoder.Decode(b)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (c Cacher[T]) GetFallback(ctx context.Context, key any, fallback cache.Fallback[T]) (T, error) {
	item, err := c.Get(ctx, key)
	if err != nil {
		if errors.Is(err, cache.ErrCacheMiss) {
			return fallback()
		}
		return item, err
	}

	return item, nil
}

func (c Cacher[T]) Del(ctx context.Context, key any) error {
	_, err := c.client.Del(ctx, c.key(key)).Result()
	return err
}
