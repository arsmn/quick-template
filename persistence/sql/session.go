package sql

import (
	"QUICK-Template/models"
	"context"

	"gorm.io/gorm"
)

type SessionPersister struct {
	GenericPersister[models.Session]
}

func newSessionPersister(db *gorm.DB) SessionPersister {
	return SessionPersister{
		GenericPersister: newGenericPersister[models.Session](db),
	}
}

func (p SessionPersister) FindByToken(ctx context.Context, token string) (models.Session, error) {
	var session models.Session
	err := p.db.Where("token = ?", token).First(&session).Error
	return session, handleError(err)
}
