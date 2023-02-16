package mongo

import (
	"QUICK-Template/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type WalletPersister struct {
	GenericPersister[models.Wallet]
}

func newWalletPersister(coll *mongo.Collection) WalletPersister {
	return WalletPersister{
		GenericPersister: newGenericPersister[models.Wallet](coll),
	}
}

func (p WalletPersister) FindByUser(ctx context.Context, wid, uid uint) (models.Wallet, error) {
	var wallet models.Wallet
	err := p.coll.FindOne(ctx, bson.M{"user_id": uid}).Decode(&wallet)
	return wallet, handleError(err)
}
