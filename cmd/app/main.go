package main

import (
	"llg_backend/config"
	"llg_backend/internal/app"
	"log"
)

func main() {
	cfg, err := config.LoadConfigFromEnv()
	if err != nil {
		log.Fatalf("cannot load config: %v", err.Error())
	}

	app.Run(cfg)
}
