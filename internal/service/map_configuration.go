package service

import (
	"context"
	"llg_backend/internal/entity"
)

type mapConfigurationService struct {
	mapConfigRepo entity.MapConfigurationRepository
	itemRepo      entity.ItemRepository
	doorRepo      entity.DoorRepository
	ruleRepo      entity.RuleRepository
}

func NewMapConfigurationService(
	mapConfigRepo entity.MapConfigurationRepository,
	itemRepo entity.ItemRepository,
	doorRepo entity.DoorRepository,
	ruleRepo entity.RuleRepository,
) entity.MapConfigurationService {
	return &mapConfigurationService{
		mapConfigRepo: mapConfigRepo,
		itemRepo:      itemRepo,
		doorRepo:      doorRepo,
		ruleRepo:      ruleRepo,
	}
}

func (s mapConfigurationService) ListFromPlayerID(ctx context.Context, playerID string) ([]*entity.PlayerStatInMap, error) {
	mapConfigs, err := s.mapConfigRepo.ListFromPlayerID(ctx, playerID)
	if err != nil {
		return nil, err
	}

	mapConfigIDs := make([]int64, 0)
	mapConfigMaps := make(map[int64]*entity.MapConfiguration)
	for _, conf := range mapConfigs {
		mapConfigIDs = append(mapConfigIDs, conf.MapConfig.ID)
		mapConfigMaps[conf.MapConfig.ID] = conf.MapConfig
	}

	mapItems, err := s.itemRepo.ListFromMapConfigurationIDs(ctx, mapConfigIDs)
	if err != nil {
		return nil, err
	}

	for _, item := range mapItems {
		if mapConfig, found := mapConfigMaps[item.MapConfigID]; found {
			mapConfig.Items = append(mapConfig.Items, item)
		}
	}

	mapDoors, err := s.doorRepo.ListFromMapConfigurationIDs(ctx, mapConfigIDs)
	if err != nil {
		return nil, err
	}

	for _, door := range mapDoors {
		if mapConfig, found := mapConfigMaps[door.MapConfigID]; found {
			mapConfig.Doors = append(mapConfig.Doors, door)
		}
	}

	mapRules, err := s.ruleRepo.ListFromMapConfigurationIDs(ctx, mapConfigIDs)
	if err != nil {
		return nil, err
	}

	for _, rule := range mapRules {
		if mapConfig, found := mapConfigMaps[rule.MapConfigID]; found {
			mapConfig.Rules = append(mapConfig.Rules, rule)
		}
	}

	return mapConfigs, nil
}
