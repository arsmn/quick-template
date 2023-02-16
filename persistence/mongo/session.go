package mongo

import (
	"QUICK-Template/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionPersister struct {
	GenericPersister[models.Session]
}

func newSessionPersister(coll *mongo.Collection) SessionPersister {
	return SessionPersister{
		GenericPersister: newGenericPersister[models.Session](coll),
	}
}

func (p SessionPersister) FindByToken(ctx context.Context, token string) (models.Session, error) {
	var session models.Session
	err := p.coll.FindOne(ctx, bson.M{"token": token}).Decode(&session)
	return session, handleError(err)
}
