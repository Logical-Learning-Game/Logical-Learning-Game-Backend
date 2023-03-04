package service

import (
	"context"
	"llg_backend/internal/entity"

	"gorm.io/gorm"
)

type mapConfigurationService struct {
	db *gorm.DB
}

func NewMapConfigurationService(db *gorm.DB) MapConfigurationService {
	return &mapConfigurationService{
		db: db,
	}
}

func (s mapConfigurationService) ListPlayerAvailableMaps(ctx context.Context, playerID string) ([]*entity.World, error) {
	var mapConfigurationForPlayers []*entity.MapConfigurationForPlayer
	result := s.db.WithContext(ctx).
		Where(&entity.MapConfigurationForPlayer{PlayerID: playerID}).
		Joins("MapConfiguration").
		Preload("MapConfiguration.Rules").
		Find(&mapConfigurationForPlayers)
	if err := result.Error; err != nil {
		return nil, err
	}

	worldIDSet := make(map[int64]struct{})
	worldIDs := make([]int64, 0, len(mapConfigurationForPlayers))
	for _, v := range mapConfigurationForPlayers {
		if _, found := worldIDSet[v.MapConfiguration.WorldID]; !found {
			worldIDs = append(worldIDs, v.MapConfiguration.WorldID)
			worldIDSet[v.MapConfiguration.WorldID] = struct{}{}
		}
	}

	var worlds []*entity.World
	result = s.db.Find(&worlds, worldIDs)
	if err := result.Error; err != nil {
		return nil, err
	}

	worldMap := make(map[int64]*entity.World)
	for _, v := range worlds {
		worldMap[v.ID] = v
	}

	for _, v := range mapConfigurationForPlayers {
		if world, found := worldMap[v.MapConfiguration.WorldID]; found {
			world.MapConfigurationForPlayers = append(world.MapConfigurationForPlayers, v)
		}
	}

	return worlds, nil
}
