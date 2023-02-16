package api

import (
	"QUICK-Template/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) usersRouter(r chi.Router) {
	r.Post("/signin", h.convert(h.signin))
	r.Post("/signup", h.convert(h.signup))
}

// Signin godoc
// @Summary      Signin user
// @Description  Authenticates user and creates a new session
// @Tags         signin
// @Accept       json
// @Produce      json
// @Param        request body models.SigninRequest true  "payload"
// @Success      200  {object}  api.PayloadResponse[models.SigninResponse]
// @Failure      400  {object}  herodot.DefaultError
// @Failure      500  {object}  herodot.DefaultError
// @Router       /users/signin [POST]
func (h *Handler) signin(rw http.ResponseWriter, r *http.Request) error {
	req, err := ParseJSON[models.SigninRequest](r)
	if err != nil {
		return err
	}

	res, err := h.service.Signin(r.Context(), req)
	if err != nil {
		return err
	}

	return h.writer.Write(rw, r, OkPayloadResponse(res))
}

// Signup godoc
// @Summary      Signup user
// @Description  Register new user
// @Tags         signup
// @Accept       json
// @Produce      json
// @Param        request body models.SignupRequest true "payload"
// @Success      200  {object}  api.Response
// @Failure      400  {object}  herodot.DefaultError
// @Failure      500  {object}  herodot.DefaultError
// @Router       /users/signup [POST]
func (h *Handler) signup(rw http.ResponseWriter, r *http.Request) error {
	req, err := ParseJSON[models.SignupRequest](r)
	if err != nil {
		return err
	}

	err = h.service.Signup(r.Context(), req)
	if err != nil {
		return err
	}

	return h.writer.Write(rw, r, OkResponse)
}
