package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GenericPersister[T any] struct {
	coll *mongo.Collection
}

func newGenericPersister[T any](coll *mongo.Collection) GenericPersister[T] {
	return GenericPersister[T]{
		coll: coll,
	}
}

func (p GenericPersister[T]) Find(ctx context.Context, id uint) (T, error) {
	var model T
	err := p.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&model)
	return model, handleError(err)
}

func (p GenericPersister[T]) Create(ctx context.Context, model T) (T, error) {
	_, err := p.coll.InsertOne(ctx, model)
	return model, handleError(err)
}

func (p GenericPersister[T]) Update(ctx context.Context, id uint, model T) (T, error) {
	_, err := p.coll.ReplaceOne(ctx, bson.M{"_id": id}, model)
	return model, handleError(err)
}

func (p GenericPersister[T]) Delete(ctx context.Context, id uint, model T) error {
	_, err := p.coll.DeleteOne(ctx, bson.M{"_id": id})
	return handleError(err)
}
