package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TransactionType string

const (
	Unknown TransactionType = ""
	Credit  TransactionType = "credit"
	Debit   TransactionType = "debit"
)

type Transaction struct {
	gorm.Model

	Amount      decimal.Decimal `gorm:"not null"`
	Type        TransactionType `gorm:"not null;size:20"`
	Description string          `gorm:"not null;size:100"`
	WalletID    uint            `gorm:"not null"`

	Wallet Wallet
}
