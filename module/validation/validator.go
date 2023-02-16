package validation

import (
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/ory/herodot"
)

func Validate(obj any) error {
	if v, ok := obj.(ozzo.Validatable); ok {
		if err := v.Validate(); err != nil {
			if es, ok := err.(ozzo.Errors); ok {
				return herodot.ErrBadRequest.WithDetail("errors", es)
			}
			return err
		}
	}

	return nil
}
