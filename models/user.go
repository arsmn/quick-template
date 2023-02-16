package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string `gorm:"not null;size:50"`
	Email    string `gorm:"not null;unique;size:50"`
	Password string `gorm:"not null;size:100"`

	Wallets []Wallet
	Session []Session
}
