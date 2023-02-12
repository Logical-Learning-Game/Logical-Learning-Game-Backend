package repository

import (
	"context"
	"llg_backend/internal/entity"
	"llg_backend/internal/entity/sqlc_generated"
)

type itemRepository struct {
	sqlc_generated.Querier
}

func NewItemRepository(querier sqlc_generated.Querier) entity.ItemRepository {
	return &itemRepository{
		Querier: querier,
	}
}

func (r *itemRepository) ListFromMapConfigurationIDs(ctx context.Context, mapConfigurationIDs []int64) ([]*entity.MapItem, error) {
	itemRows, err := r.Querier.GetItemFromMapConfigurationIDs(ctx, mapConfigurationIDs)
	if err != nil {
		return nil, err
	}

	mapItems := make([]*entity.MapItem, 0)
	for _, row := range itemRows {
		mapItems = append(mapItems, &entity.MapItem{
			MapConfigID: row.MapConfigurationID,
			Name:        row.Name,
			Type:        row.Type,
			Position: entity.Vector2Int{
				X: int(row.PositionX),
				Y: int(row.PositionY),
			},
		})
	}

	return mapItems, nil
}
