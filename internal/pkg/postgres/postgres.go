package postgres

import (
	"llg_backend/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	URI string `mapstructure:"POSTGRES_URI"`
}

func New(config *Config) (*gorm.DB, error) {
	logger.GlobalLog.Infof("connecting to postgres at %s", config.URI)

	db, err := gorm.Open(postgres.Open(config.URI), &gorm.Config{})
	if err != nil {
		logger.GlobalLog.Errorw("connect to postgres failed", "err", err)
		return nil, err
	}

	logger.GlobalLog.Infof("connecting to postgres successful")
	return db, nil
}
