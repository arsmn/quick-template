package sql

import (
	"QUICK-Template/models"
	"context"

	"gorm.io/gorm"
)

type WalletPersister struct {
	GenericPersister[models.Wallet]
}

func newWalletPersister(db *gorm.DB) WalletPersister {
	return WalletPersister{
		GenericPersister: newGenericPersister[models.Wallet](db),
	}
}

func (p WalletPersister) FindByUser(ctx context.Context, wid, uid uint) (models.Wallet, error) {
	var wallet models.Wallet
	err := p.db.Where("user_id = ?", uid).First(&wallet, wid).Error
	return wallet, handleError(err)
}
