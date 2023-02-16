package sql

import (
	"QUICK-Template/models"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionPersister struct {
	GenericPersister[models.Transaction]
}

func newTransactionPersister(db *gorm.DB) TransactionPersister {
	return TransactionPersister{
		GenericPersister: newGenericPersister[models.Transaction](db),
	}
}

func (p TransactionPersister) create(ctx context.Context, transaction models.Transaction, updateBalance clause.Expr) error {
	err := p.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		if err := tx.Model(&transaction.Wallet).
			Update("balance", updateBalance).Error; err != nil {
			return err
		}

		return nil
	})

	return handleError(err)
}

func (p TransactionPersister) Credit(ctx context.Context, transaction models.Transaction) error {
	return p.create(ctx, transaction, gorm.Expr("balance + ?", transaction.Amount))
}

func (p TransactionPersister) Debit(ctx context.Context, transaction models.Transaction) error {
	return p.create(ctx, transaction, gorm.Expr("balance - ?", transaction.Amount))
}
