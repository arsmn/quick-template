package mongo

import (
	"context"

	"github.com/cenkalti/backoff"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	DSN string
}

func Connect(cfg Config) (*mongo.Client, error) {
	var client *mongo.Client

	operration := func() (err error) {
		client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.DSN))
		if err != nil {
			return handleError(err)
		}

		if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
			return handleError(err)
		}

		return nil
	}

	return client,
		backoff.Retry(operration, backoff.NewExponentialBackOff())
}
