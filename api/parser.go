package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ory/herodot"
)

func ParseJSON[T any](r *http.Request) (T, error) {
	var input T

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return input, herodot.ErrInternalServerError.WithWrap(err)
	}

	if err := r.Body.Close(); err != nil {
		return input, herodot.ErrInternalServerError.WithWrap(err)
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	if err := json.Unmarshal(body, &input); err != nil {
		return input, herodot.ErrInternalServerError.WithWrap(err)
	}

	return input, nil
}
