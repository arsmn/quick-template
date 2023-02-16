package mongo

import (
	"errors"

	"github.com/ory/herodot"
	"go.mongodb.org/mongo-driver/mongo"
)

func handleError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		return herodot.ErrNotFound
	}

	return err
}
