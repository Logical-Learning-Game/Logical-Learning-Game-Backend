package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"llg_backend/internal/pkg/postgres"
)

type Config struct {
	HTTP      HTTP            `mapstructure:",squash"`
	Postgres  postgres.Config `mapstructure:",squash"`
}

type HTTP struct {
	Port string `mapstructure:"SERVER_PORT"`
}

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
