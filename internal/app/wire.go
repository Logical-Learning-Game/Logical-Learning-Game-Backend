//go:build wireinject

package app

import (
	v1 "llg_backend/internal/controller/http/v1"
	"llg_backend/internal/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var providerSet = wire.NewSet(
	service.NewPlayerStatisticService,
	service.NewMapConfigurationService,
)

func InitializePlayerController(db *gorm.DB) *v1.PlayerController {
	wire.Build(
		v1.NewPlayerController,
		providerSet,
	)
	return &v1.PlayerController{}
}
