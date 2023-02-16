package api

import (
	"QUICK-Template/models"
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ory/herodot"
)

var (
	OkResponse = Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	}
)

type Response struct {
	Code   int    `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
}

type PayloadResponse[T any] struct {
	Response
	Data T `json:"data,omitempty"`
}

func OkPayloadResponse[T any](data T) PayloadResponse[T] {
	return PayloadResponse[T]{
		Data:     data,
		Response: OkResponse,
	}
}

type ctxKey int

const (
	unknownKey ctxKey = iota
	sessionKey
)

func contextParam[T any](ctx context.Context, key ctxKey) (T, error) {
	p, ok := ctx.Value(key).(T)
	if !ok {
		return p, errors.New("key not found")
	}
	return p, nil
}

func mustContextParam[T any](ctx context.Context, key ctxKey) T {
	p, err := contextParam[T](ctx, key)
	if err != nil {
		panic(err)
	}
	return p
}

func sessionParam(ctx context.Context) models.Session {
	return mustContextParam[models.Session](ctx, sessionKey)
}

func authorizedRequest(ctx context.Context) models.AuthorizedRequest {
	return models.AuthorizedRequest{
		Session: sessionParam(ctx),
	}
}

func urlParamUint(r *http.Request, key string) (uint, error) {
	p, err := strconv.ParseUint(chi.URLParam(r, key), 10, 0)
	if err != nil {
		return 0, err
	}

	return uint(p), nil
}

func walletIDParam(r *http.Request) (uint, error) {
	wid, err := urlParamUint(r, WalletIDParam)
	if err != nil {
		return 0, herodot.ErrBadRequest.WithReason("invalid wallet id parameter")
	}

	return wid, nil
}
