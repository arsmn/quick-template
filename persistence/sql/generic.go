package sql

import (
	"context"

	"gorm.io/gorm"
)

type GenericPersister[T any] struct {
	db *gorm.DB
}

func newGenericPersister[T any](db *gorm.DB) GenericPersister[T] {
	return GenericPersister[T]{
		db: db,
	}
}

func (p GenericPersister[T]) Find(ctx context.Context, id uint) (T, error) {
	var model T
	err := p.db.First(&model, id).Error
	return model, handleError(err)
}

func (p GenericPersister[T]) Create(ctx context.Context, model T) (T, error) {
	res := p.db.Create(&model)
	return model, handleError(res.Error)
}

func (p GenericPersister[T]) Update(ctx context.Context, id uint, model T) (T, error) {
	res := p.db.Save(&model)
	return model, handleError(res.Error)
}

func (p GenericPersister[T]) Delete(ctx context.Context, id uint, model T) error {
	res := p.db.Delete(&model)
	return handleError(res.Error)
}
