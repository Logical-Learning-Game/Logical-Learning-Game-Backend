package repository

import (
	"context"
	"llg_backend/internal/entity"
	"llg_backend/internal/entity/sqlc_generated"
)

type worldRepository struct {
	sqlc_generated.Querier
}

func NewWorldRepository(querier sqlc_generated.Querier) entity.WorldRepository {
	return &worldRepository{
		Querier: querier,
	}
}

func (r *worldRepository) ListFromMapConfigurationIDs(ctx context.Context, mapConfigurationIDs []int64) ([]*entity.World, error) {
	worldRows, err := r.Querier.ListWorldFromMapConfigurationIDs(ctx, mapConfigurationIDs)
	if err != nil {
		return nil, err
	}

	playerWorlds := make([]*entity.World, 0)
	for _, row := range worldRows {
		playerWorlds = append(playerWorlds, &entity.World{
			ID:   row.WorldID,
			Name: row.Name,
			Maps: make([]*entity.PlayerStatInMap, 0),
		})
	}

	return playerWorlds, nil
}
