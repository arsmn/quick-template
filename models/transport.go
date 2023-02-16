package models

import (
	"QUICK-Template/module/validation"
	"time"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/shopspring/decimal"
)

type AuthorizedRequest struct {
	Session `json:"-" swaggerignore:"true"`
}

type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req SigninRequest) Validate() error {
	return v.ValidateStruct(&req,
		v.Field(&req.Email, v.Required),
		v.Field(&req.Password, v.Required),
	)
}

type SigninResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req SignupRequest) Validate() error {
	return v.ValidateStruct(&req,
		v.Field(&req.Name, v.Required, v.Length(3, 50)),
		v.Field(&req.Email, v.Required, is.Email),
		v.Field(&req.Password, v.Required, v.Length(6, 50)),
	)
}

type GetSessionRequest struct {
	Token string `json:"-"`
}

type GetSessionResponse struct {
	Session Session `json:"-"`
}

type GetWalletRequest struct {
	AuthorizedRequest
	WalletID uint `json:"-"`
}

type GetWalletResponse struct {
	ID      uint            `json:"id"`
	Balance decimal.Decimal `json:"balance"`
	UserID  uint            `json:"user_id"`
}

type CreditWalletRequest struct {
	AuthorizedRequest
	WalletID    uint            `json:"-"`
	Amount      decimal.Decimal `json:"amount"`
	Description string          `json:"description"`
}

func (req CreditWalletRequest) Validate() error {
	return v.ValidateStruct(&req,
		v.Field(&req.Description, v.Length(3, 100)),
		v.Field(&req.Amount, validation.MinDecimal(decimal.Zero)),
	)
}

type DebitWalletRequest struct {
	AuthorizedRequest
	WalletID    uint            `json:"-"`
	Amount      decimal.Decimal `json:"amount"`
	Description string          `json:"description"`
}

func (req DebitWalletRequest) Validate() error {
	return v.ValidateStruct(&req,
		v.Field(&req.Description, v.Length(3, 100)),
		v.Field(&req.Amount, validation.MinDecimal(decimal.Zero)),
	)
}
