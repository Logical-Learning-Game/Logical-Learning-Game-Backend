package main

import (
	"llg_backend/config"
	"llg_backend/internal/app"
	"log"

	"github.com/spf13/viper"
)

func main() {
	}

	// Read secret environment variable
	if err := config.LoadConfigEnv(); err != nil {
		log.Fatalf("cannot load env config: %v", err.Error())
	}

	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("cannot unmarshal config: %v", err.Error())
	}

	app.Run(&cfg)
}
