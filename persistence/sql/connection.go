package sql

import (
	"QUICK-Template/models"

	"github.com/cenkalti/backoff/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DSN string
}

func Connect(cfg Config) (*gorm.DB, error) {
	var db *gorm.DB

	operration := func() (err error) {
		db, err = gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
		if err != nil {
			return handleError(err)
		}

		if err := db.AutoMigrate(
			&models.User{},
			&models.Wallet{},
			&models.Transaction{},
			&models.Session{},
		); err != nil {
			return handleError(err)
		}

		return nil
	}

	return db,
		backoff.Retry(operration, backoff.NewExponentialBackOff())
}
