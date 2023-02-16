package sql

import (
	"QUICK-Template/models"
	"context"

	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB

	user        models.UserPersister
	wallet      models.WalletPersister
	transaction models.TransactionPersister
	session     models.SessionPersister
}

func NewStorage(db *gorm.DB) Storage {
	return Storage{
		db:          db,
		user:        newUserPersister(db),
		wallet:      newWalletPersister(db),
		transaction: newTransactionPersister(db),
		session:     newSessionPersister(db),
	}
}

func (s Storage) User() models.UserPersister {
	return s.user
}

func (s Storage) Wallet() models.WalletPersister {
	return s.wallet
}

func (s Storage) Transaction() models.TransactionPersister {
	return s.transaction
}

func (s Storage) Session() models.SessionPersister {
	return s.session
}

func (s Storage) Close(ctx context.Context) error {
	sqldb, err := s.db.DB()
	if err != nil {
		return err
	}

	return sqldb.Close()
}
