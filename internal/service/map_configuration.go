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

func (s mapConfigurationService) ListFromPlayerID(ctx context.Context, playerID string) ([]*entity.PlayerMapConfiguration, error) {
	mapConfigs, err := s.mapConfigRepo.ListFromPlayerID(ctx, playerID)
	if err != nil {
		return nil, err
	}

	mapConfigIDs := make([]int64, 0)
	mapConfigMaps := make(map[int64]*entity.PlayerMapConfiguration)
	for _, conf := range mapConfigs {
		mapConfigIDs = append(mapConfigIDs, conf.MapConfiguration.ID)
		mapConfigMaps[conf.MapConfiguration.ID] = conf
	}

	mapItems, err := s.itemRepo.ListFromMapConfigurationIDs(ctx, mapConfigIDs)
	if err != nil {
		return nil, err
	}

	for _, item := range mapItems {
		if playerMap, found := mapConfigMaps[item.MapConfigID]; found {
			playerMap.MapConfiguration.Items = append(playerMap.MapConfiguration.Items, item)
		}
	}

	mapDoors, err := s.doorRepo.ListFromMapConfigurationIDs(ctx, mapConfigIDs)
	if err != nil {
		return nil, err
	}

	for _, door := range mapDoors {
		if playerMap, found := mapConfigMaps[door.MapConfigID]; found {
			playerMap.MapConfiguration.Doors = append(playerMap.MapConfiguration.Doors, door)
		}
	}

	mapRules, err := s.ruleRepo.ListFromMapConfigurationIDs(ctx, mapConfigIDs)
	if err != nil {
		return nil, err
	}

	for _, rule := range mapRules {
		if playerMap, found := mapConfigMaps[rule.MapConfigID]; found {
			playerMap.MapConfiguration.Rules = append(playerMap.MapConfiguration.Rules, rule)
		}
	}

	return mapConfigs, nil
}
