package sql

import (
	"errors"

	"github.com/ory/herodot"
	"gorm.io/gorm"
)

func handleError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return herodot.ErrNotFound
	}

	return err
}
