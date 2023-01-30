package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	HTTP      `mapstructure:",squash"`
	Postgres  `mapstructure:",squash"`
}

type HTTP struct {
	Port string `mapstructure:"SERVER_PORT"`
}

type Postgres struct {
	URI string `mapstructure:"POSTGRES_URI"`
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
