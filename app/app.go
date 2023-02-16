package app

import (
	"QUICK-Template/cache"
	"QUICK-Template/logger"
	"QUICK-Template/models"
	"QUICK-Template/module/hash"
)

type App struct {
	logger       *logger.Logger
	storage      models.Storage
	hasher       hash.Hasher
	walletCache  cache.Cacher[models.Wallet]
	sessionCache cache.Cacher[models.Session]
}

func New(l *logger.Logger,
	s models.Storage,
	h hash.Hasher,
	wc cache.Cacher[models.Wallet],
	sc cache.Cacher[models.Session]) App {
	return App{
		logger:       l,
		storage:      s,
		hasher:       h,
		walletCache:  wc,
		sessionCache: sc,
	}
}
