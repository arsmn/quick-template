package models

import (
	"QUICK-Template/module/generate"
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model

	Token     string    `gorm:"not null;unique;size:100"`
	ExpiresAt time.Time `gorm:"not null"`
	UserID    uint      `gorm:"not null"`

	User User
}

func NewSession(uid uint, ttl time.Duration) (Session, error) {
	var s Session

	token, err := generate.UUID()
	if err != nil {
		return s, err
	}

	s.UserID = uid
	s.Token = token
	s.ExpiresAt = time.Now().UTC().Add(ttl)

	return s, nil
}

func (s Session) IsExpired(base time.Time) bool {
	return base.After(s.ExpiresAt)
}

func (s Session) IsValid() bool {
	return !s.IsExpired(time.Now())
}
