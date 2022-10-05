package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	HTTP    `mapstructure:",squash"`
	MariaDB `mapstructure:",squash"`
}

type HTTP struct {
	Port string `mapstructure:"SERVER_PORT"`
}

type MariaDB struct {
	DBSource string `mapstructure:"DB_SOURCE"`
}

func LoadConfigFromEnv() (*Config, error) {
	viper.AutomaticEnv()

	var config Config

	envKeysMap := new(map[string]interface{})
	if err := mapstructure.Decode(config, envKeysMap); err != nil {
		return nil, err
	}

	for k := range *envKeysMap {
		if err := viper.BindEnv(k); err != nil {
			return nil, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config

	err := viper.Unmarshal(&config)
	return &config, err
}
