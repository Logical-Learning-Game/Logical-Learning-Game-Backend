package postgres

import (
	"fmt"
	"llg_backend/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     string `mapstructure:"POSTGRES_PORT"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	DBName   string `mapstructure:"POSTGRES_DB"`
	SSLMode  string `mapstructure:"POSTGRES_SSL_MODE"`
	Debug    bool   `mapstructure:"GORM_DEBUG"`
}

func New(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)

	logger.GlobalLog.Infof("connecting to postgres at %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.GlobalLog.Errorw("connect to postgres failed", "err", err)
		return nil, err
	}

	if config.Debug {
		db = db.Debug()
	}

	logger.GlobalLog.Infof("connecting to postgres successful")
	return db, nil
}
