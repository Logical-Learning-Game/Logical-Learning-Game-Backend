//go:build wireinject

package v1

import (
	"llg_backend/internal/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var providerSet = wire.NewSet(
	service.NewPlayerStatisticService,
	service.NewMapConfigurationService,
	service.NewPlayerService,
)

func InitializePlayerController(db *gorm.DB) *PlayerController {
	wire.Build(
		NewPlayerController,
		providerSet,
	)
	return &PlayerController{}
}
