package user

import (
	"QUICK-Template/models"
	"context"
)

type Client interface {
	GetUser(ctx context.Context, uid int64) (*models.User, error)
}
