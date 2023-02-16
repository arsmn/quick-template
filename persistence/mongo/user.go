package mongo

import (
	"QUICK-Template/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserPersister struct {
	GenericPersister[models.User]
}

func newUserPersister(coll *mongo.Collection) UserPersister {
	return UserPersister{
		GenericPersister: newGenericPersister[models.User](coll),
	}
}

func (p UserPersister) FindByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := p.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return user, handleError(err)
}
