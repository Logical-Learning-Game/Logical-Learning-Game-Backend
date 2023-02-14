package service

import (
	"context"
	"llg_backend/internal/entity"
)

type worldService struct {
	mapConfigService entity.MapConfigurationService
	worldRepo        entity.WorldRepository
}

func NewWorldService(mapConfigService entity.MapConfigurationService, worldRepo entity.WorldRepository) entity.WorldService {
	return &worldService{
		mapConfigService: mapConfigService,
		worldRepo:        worldRepo,
	}
}

func (s worldService) ListFromPlayerID(ctx context.Context, playerID string) ([]*entity.World, error) {
	mapConfigs, err := s.mapConfigService.ListFromPlayerID(ctx, playerID)
	if err != nil {
		return nil, err
	}

	mapConfigIDs := make([]int64, 0)
	for _, conf := range mapConfigs {
		mapConfigIDs = append(mapConfigIDs, conf.MapConfiguration.ID)
	}

	worlds, err := s.worldRepo.ListFromMapConfigurationIDs(ctx, mapConfigIDs)
	if err != nil {
		return nil, err
	}

	worldMap := make(map[int64]*entity.World)
	for _, world := range worlds {
		worldMap[world.ID] = world
	}

	for _, conf := range mapConfigs {
		if world, found := worldMap[conf.MapConfiguration.WorldID]; found {
			world.Maps = append(world.Maps, conf)
		}
	}

	return worlds, nil
}
