package service

import (
	"context"
	"llg_backend/internal/dto"
	"llg_backend/internal/dto/mapper"
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

func (s mapConfigurationService) ListPlayerAvailableMaps(ctx context.Context, playerID string) ([]*dto.WorldDTO, error) {
	var mapConfigurations []*entity.MapConfiguration
	result := s.db.WithContext(ctx).
		Table("map_configurations AS map").
		Joins("INNER JOIN map_configuration_for_players AS map_player ON map_player.map_configuration_id = map.id").
		Where("map_player.player_id = ?", playerID).
		Order("map.id ASC").
		Preload("Rules").
		Find(&mapConfigurations)
	if err := result.Error; err != nil {
		return nil, err
	}

	worldIDSet := make(map[int64]struct{})
	worldIDs := make([]int64, 0, len(mapConfigurations))
	for _, v := range mapConfigurations {
		if _, found := worldIDSet[v.WorldID]; !found {
			worldIDs = append(worldIDs, v.WorldID)
			worldIDSet[v.WorldID] = struct{}{}
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

	for _, v := range mapConfigurations {
		if world, found := worldMap[v.WorldID]; found {
			world.MapConfigurations = append(world.MapConfigurations, v)
		}
	}

	worldMapper := mapper.NewWorldMapper()
	worldDTOs := make([]*dto.WorldDTO, 0, len(worlds))
	for _, world := range worlds {
		worldDTO := worldMapper.ToDTO(world)
		worldDTOs = append(worldDTOs, worldDTO)
	}

	return worldDTOs, nil
}
