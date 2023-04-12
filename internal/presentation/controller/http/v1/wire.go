//go:build wireinject

package v1

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"llg_backend/config"
	"llg_backend/internal/service"
	"llg_backend/internal/token"
)

var providerSet = wire.NewSet(
	service.NewPlayerStatisticService,
	service.NewMapConfigurationService,
	service.NewPlayerService,
	service.NewAdminAuthenticationService,
)

func InitializePlayerController(db *gorm.DB) *PlayerController {
	wire.Build(
		NewPlayerController,
		providerSet,
	)
	return &PlayerController{}
}

func InitializeAdminController(db *gorm.DB) *AdminController {
	wire.Build(
		NewAdminController,
		providerSet,
	)
	return &AdminController{}
}

func InitializeAdminAuthenticationController(cfg *config.Config, db *gorm.DB, tokenMaker token.Maker) *AdminAuthenticationController {
	wire.Build(
		NewAdminAuthenticationController,
		providerSet,
	)
	return &AdminAuthenticationController{}
}
