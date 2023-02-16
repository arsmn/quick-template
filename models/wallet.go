package models

import (
	"github.com/ory/herodot"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
)

type Wallet struct {
	gorm.Model

	Balance decimal.Decimal        `gorm:"not null"`
	UserID  uint                   `gorm:"not null"`
	Version optimisticlock.Version `gorm:"not null"`

	User        User
	Transaction []Transaction
}

func (w Wallet) IsOwner(uid uint) bool {
	return w.UserID == uid
}

func (w *Wallet) Credit(amount decimal.Decimal) error {
	if amount.LessThan(decimal.Zero) {
		return herodot.ErrBadRequest.WithReason("amount must be greater than 0.")
	}

	w.Balance = w.Balance.Add(amount)

	return nil
}

func (w *Wallet) Debit(amount decimal.Decimal) error {
	if amount.LessThan(decimal.Zero) {
		return herodot.ErrBadRequest.WithReason("amount must be greater than 0.")
	}

	balance := w.Balance.Sub(amount)
	if balance.LessThan(decimal.Zero) {
		return herodot.ErrBadRequest.WithReason("credit is not enough")
	}

	w.Balance = balance

	return nil
}
