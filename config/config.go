package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"llg_backend/internal/pkg/postgres"
	"time"
)

type Config struct {
	HTTP     HTTP            `mapstructure:",squash"`
	Postgres postgres.Config `mapstructure:",squash"`
	JWT      JWT             `mapstructure:",squash"`
}

type HTTP struct {
	Port string `mapstructure:"SERVER_PORT"`
}

type JWT struct {
	SecretKey string        `mapstructure:"JWT_SECRET_KEY"`
	Duration  time.Duration `mapstructure:"JWT_DURATION"`
}

func LoadConfigEnv() error {
	var config Config

	envKeysMap := new(map[string]interface{})
	if err := mapstructure.Decode(config, envKeysMap); err != nil {
		return err
	}

	for k := range *envKeysMap {
		if err := viper.BindEnv(k); err != nil {
			return err
		}
	}

	return nil
}
