package app

import "github.com/ory/herodot"

var (
	ErrPasswordMismatch = herodot.ErrBadRequest.WithReason("email and password do not match")
)
