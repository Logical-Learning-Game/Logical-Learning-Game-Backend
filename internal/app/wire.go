//go:build wireinject

package app

import (
	v1 "llg_backend/internal/controller/http/v1"
	"llg_backend/internal/entity/sqlc_generated"
	"llg_backend/internal/repository"
	"llg_backend/internal/service"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	service.NewWorldService,
	service.NewPlayerService,
	repository.NewPlayerRepository,
	service.NewMapConfigurationService,
	repository.NewItemRepository,
	repository.NewDoorRepository,
	repository.NewRuleRepository,
	repository.NewMapConfigurationRepository,
	repository.NewWorldRepository,
)

func InitializePlayerController(querier sqlc_generated.Querier) *v1.PlayerController {
	wire.Build(
		v1.NewPlayerController,
		providerSet,
	)
	return &v1.PlayerController{}
}
