package api

import (
	"QUICK-Template/logger"
	"QUICK-Template/models"
	"context"
	"errors"
	"net/http"

	"github.com/ory/herodot"
)

func (h *Handler) AuthorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if len(token) == 0 {
			h.writer.WriteError(rw, r, herodot.ErrUnauthorized)
			return
		}

		res, err := h.service.Session(r.Context(), models.GetSessionRequest{Token: token})
		if err != nil {
			if errors.Is(err, herodot.ErrNotFound) {
				err = herodot.ErrUnauthorized
			}

			h.writer.WriteError(rw, r, err)
			return
		}

		if !res.Session.IsValid() {
			h.writer.WriteError(rw, r, herodot.ErrUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), sessionKey, res.Session)

		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func (h *Handler) LoggerMiddleware(logger *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			logger.Info("Method", r.Method)
			next.ServeHTTP(rw, r)
			logger.Info("Status", r.Response.Status)
		})
	}
}
