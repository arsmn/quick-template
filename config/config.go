package config

import (
	"QUICK-Template/api"
	"QUICK-Template/cache/redis"
	"QUICK-Template/persistence/sql"

	"github.com/spf13/viper"
)

const (
	keySQLDSN        = "SQL.DSN"
	keyAPIPort       = "API.Port"
	keyRedisAddr     = "REDIS.ADDR"
	keyRedisPassword = "REDIS.PASSWORD"
)

type Config struct {
	SQL   sql.Config
	API   api.Config
	Redis redis.Config
}

func New() Config {
	var cfg Config

	cfg.SQL.DSN = viper.GetString(keySQLDSN)
	cfg.API.Port = viper.GetString(keyAPIPort)
	cfg.Redis.Address = viper.GetString(keyRedisAddr)
	cfg.Redis.Password = viper.GetString(keyRedisPassword)

	return cfg
}
