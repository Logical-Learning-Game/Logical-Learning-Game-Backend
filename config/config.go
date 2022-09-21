package config

import "github.com/spf13/viper"

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

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	var config Config
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	err := viper.Unmarshal(&config)

	return &config, err
}
