//go:build wireinject

package app

import (
	"database/sql"
	v1 "llg_backend/internal/controller/http/v1"
	"llg_backend/internal/entity/sqlc_generated"
	"llg_backend/internal/repository"
	"llg_backend/internal/service"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	service.NewPlayerStatisticService,
	service.NewWorldService,
	service.NewPlayerService,
	repository.NewPlayerRepository,
	service.NewMapConfigurationService,
	repository.NewItemRepository,
	repository.NewDoorRepository,
	repository.NewRuleRepository,
	repository.NewMapConfigurationRepository,
	repository.NewWorldRepository,
	repository.NewPlayHistoryRepository,
	repository.NewGameSessionRepository,
	repository.NewUnitOfWork,
)

func InitializePlayerController(querier sqlc_generated.Querier, db *sql.DB) *v1.PlayerController {
	wire.Build(
		v1.NewPlayerController,
		providerSet,
	)
	return &v1.PlayerController{}
}
