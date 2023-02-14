package repository

import (
	"context"
	"llg_backend/internal/entity"
	"llg_backend/internal/entity/sqlc_generated"
	"llg_backend/pkg/utility"
)

type mapConfigurationRepository struct {
	sqlc_generated.Querier
}

func NewMapConfigurationRepository(querier sqlc_generated.Querier) entity.MapConfigurationRepository {
	return &mapConfigurationRepository{
		Querier: querier,
	}
}

func (r mapConfigurationRepository) ListFromPlayerID(ctx context.Context, playerID string) ([]*entity.PlayerMapConfiguration, error) {
	mapConfigRows, err := r.Querier.GetMapConfigFromPlayerID(ctx, playerID)
	if err != nil {
		return nil, err
	}

	playerMaps := make([]*entity.PlayerMapConfiguration, 0)
	for _, row := range mapConfigRows {

		intConvertionSlice := make([]int, 0)
		for _, val := range row.TileArray {
			intConvertionSlice = append(intConvertionSlice, int(val))
		}

		mapHeight := int(row.MapHeight)
		mapWidth := int(row.MapWidth)
		twoDimensionConvertionSlice := utility.TwoDimensionSlice[int](intConvertionSlice, mapHeight, mapWidth)

		mapConfiguration := &entity.MapConfiguration{
			BadgeRequirement: entity.BadgeRequirement{
				LeastSolvableCommandGold:   int(row.LeastSolvableCommandGold),
				LeastSolvableCommandSilver: int(row.LeastSolvableCommandSilver),
				LeastSolvableCommandBronze: int(row.LeastSolvableCommandBronze),
				LeastSolvableActionGold:    int(row.LeastSolvableActionGold),
				LeastSolvableActionSilver:  int(row.LeastSolvableActionSilver),
				LeastSolvableActionBronze:  int(row.LeastSolvableActionBronze),
			},
			ID:                   row.MapConfigID,
			WorldID:              row.WorldID,
			ConfigName:           row.ConfigName,
			Map:                  twoDimensionConvertionSlice,
			Height:               mapHeight,
			Width:                mapWidth,
			StartPlayerDirection: row.StartPlayerDirection,
			StartPlayerPosition: entity.Vector2Int{
				X: int(row.StartPlayerPositionX),
				Y: int(row.StartPlayerPositionY),
			},
			GoalPosition: entity.Vector2Int{
				X: int(row.GoalPositionX),
				Y: int(row.GoalPositionY),
			},
			MapImagePath:    row.MapImagePath.String,
			Difficulty:      row.Difficulty,
			StarRequirement: int(row.StarRequirement),
			Items:           make([]*entity.MapItem, 0),
			Doors:           make([]*entity.MapDoor, 0),
			Rules:           make([]*entity.MapRule, 0),
		}

		playerMapConfiguration := &entity.PlayerMapConfiguration{
			MapConfiguration: mapConfiguration,
			IsPass:           row.IsPass,
			TopSubmitHistory: nil,
		}

		playerMaps = append(playerMaps, playerMapConfiguration)
	}

	return playerMaps, nil
}
