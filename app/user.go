package app

import (
	"QUICK-Template/models"
	"QUICK-Template/module/hash"
	v "QUICK-Template/module/validation"
	"context"
	"errors"
	"time"

	"github.com/ory/herodot"
)

func (a App) Signin(ctx context.Context, req models.SigninRequest) (models.SigninResponse, error) {
	if err := v.Validate(req); err != nil {
		return models.SigninResponse{}, err
	}

	user, err := a.storage.User().FindByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, herodot.ErrNotFound) {
			return models.SigninResponse{}, ErrPasswordMismatch
		}
		return models.SigninResponse{}, err
	}

	err = a.hasher.Compare(ctx, []byte(req.Password), []byte(user.Password))
	if err != nil {
		if errors.Is(err, hash.ErrMismatchedHashAndData) {
			return models.SigninResponse{}, ErrPasswordMismatch
		}
		return models.SigninResponse{}, err
	}

	sess, err := models.NewSession(user.ID, 12*time.Hour)
	if err != nil {
		return models.SigninResponse{}, err
	}

	_, err = a.storage.Session().Create(ctx, sess)
	if err != nil {
		return models.SigninResponse{}, err
	}

	return models.SigninResponse{
		Token:     sess.Token,
		ExpiresAt: sess.ExpiresAt,
	}, nil
}

func (a App) Signup(ctx context.Context, req models.SignupRequest) error {
	if err := v.Validate(req); err != nil {
		return err
	}

	_, err := a.storage.User().FindByEmail(ctx, req.Email)
	if err == nil {
		return herodot.ErrBadRequest.WithReason("email is already registered")
	} else if !errors.Is(err, herodot.ErrNotFound) {
		return err
	}

	pswd, err := a.hasher.Hash(ctx, []byte(req.Password))
	if err != nil {
		return err
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(pswd),
	}

	_, err = a.storage.User().Create(ctx, user)

	return err
}

func (a App) Session(ctx context.Context, req models.GetSessionRequest) (models.GetSessionResponse, error) {
	fallback := func() (models.Session, error) {
		res, err := a.storage.Session().FindByToken(ctx, req.Token)
		if err != nil {
			return res, err
		}

		if err := a.sessionCache.Set(ctx, req.Token, res); err != nil {
			a.logger.Error(err)
		}

		return res, nil
	}

	session, err := a.sessionCache.GetFallback(ctx, req.Token, fallback)
	if err != nil {
		return models.GetSessionResponse{}, err
	}

	return models.GetSessionResponse{Session: session}, nil
}
