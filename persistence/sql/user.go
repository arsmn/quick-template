package sql

import (
	"QUICK-Template/models"
	"context"

	"gorm.io/gorm"
)

type UserPersister struct {
	GenericPersister[models.User]
}

func newUserPersister(db *gorm.DB) UserPersister {
	return UserPersister{
		GenericPersister: newGenericPersister[models.User](db),
	}
}

func (p UserPersister) FindByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := p.db.Where("email = ?", email).First(&user).Error
	return user, handleError(err)
}
