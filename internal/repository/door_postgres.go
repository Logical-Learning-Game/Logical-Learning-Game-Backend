package repository

import (
	"context"
	"llg_backend/internal/entity"
	"llg_backend/internal/entity/sqlc_generated"
)

type doorRepository struct {
	sqlc_generated.Querier
}

func NewDoorRepository(querier sqlc_generated.Querier) entity.DoorRepository {
	return &doorRepository{
		Querier: querier,
	}
}

func (r *doorRepository) ListFromMapConfigurationIDs(ctx context.Context, mapConfigurationIDs []int64) ([]*entity.MapDoor, error) {
	doorRows, err := r.Querier.GetDoorFromMapConfigurationIDs(ctx, mapConfigurationIDs)
	if err != nil {
		return nil, err
	}

	mapDoors := make([]*entity.MapDoor, 0)
	for _, row := range doorRows {
		mapDoors = append(mapDoors, &entity.MapDoor{
			MapConfigID: row.MapConfigurationID,
			Name:        row.Name,
			Type:        row.Type,
			Position: entity.Vector2Int{
				X: int(row.PositionX),
				Y: int(row.PositionY),
			},
			DoorDirection: row.DoorDirection,
		})
	}

	return mapDoors, nil
}
