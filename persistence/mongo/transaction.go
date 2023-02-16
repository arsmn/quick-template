package mongo

import (
	"QUICK-Template/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionPersister struct {
	cli *mongo.Client
	GenericPersister[models.Transaction]
}

func newTransactionPersister(cli *mongo.Client, coll *mongo.Collection) TransactionPersister {
	return TransactionPersister{
		GenericPersister: newGenericPersister[models.Transaction](coll),
	}
}

func (p TransactionPersister) create(ctx context.Context, transaction models.Transaction, updateBalance bson.M) error {
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		if _, err := p.coll.InsertOne(sessCtx, transaction); err != nil {
			return nil, err
		}

		if _, err := p.cli.Database("").Collection("").UpdateOne(sessCtx, bson.M{"_id": transaction.WalletID}, updateBalance); err != nil {
			return nil, err
		}
		return nil, nil
	}

	session, err := p.cli.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	if _, err := session.WithTransaction(ctx, callback); err != nil {
		return err
	}

	return handleError(err)
}

func (p TransactionPersister) Credit(ctx context.Context, transaction models.Transaction) error {
	return p.create(ctx, transaction, bson.M{})
}

func (p TransactionPersister) Debit(ctx context.Context, transaction models.Transaction) error {
	return p.create(ctx, transaction, bson.M{})
}
