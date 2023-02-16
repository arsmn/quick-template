package redis

import (
	"context"

	"github.com/cenkalti/backoff/v4"
	"github.com/go-redis/redis/v9"
)

type Config struct {
	Address  string
	Password string
	DB       int
}

func Connect(cfg Config) (*redis.Client, error) {
	var client *redis.Client

	operration := func() error {
		client = redis.NewClient(&redis.Options{
			Addr:     cfg.Address,
			Password: cfg.Password,
			DB:       cfg.DB,
		})

		if err := client.Ping(context.Background()).Err(); err != nil {
			return err
		}

		return nil
	}

	return client,
		backoff.Retry(operration, backoff.NewExponentialBackOff())
}
