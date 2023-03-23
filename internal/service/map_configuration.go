package service

import (
	"context"
	"gorm.io/gorm"
	"llg_backend/internal/dto"
	"llg_backend/internal/dto/mapper"
	"llg_backend/internal/entity"
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
		Where("map_player.player_id = ? AND active = true", playerID).
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
		worldDTO := worldMapper.ToWorldDTO(world)
		worldDTOs = append(worldDTOs, worldDTO)
	}

	return worldDTOs, nil
}

func (s mapConfigurationService) ListWorld(ctx context.Context) ([]*dto.WorldForAdminResponse, error) {
	worlds := make([]*entity.World, 0)

	result := s.db.WithContext(ctx).
		Order("id ASC").
		Find(&worlds)
	if err := result.Error; err != nil {
		return nil, err
	}

	worldsForAdmin := make([]*dto.WorldForAdminResponse, 0, len(worlds))
	for _, v := range worlds {
		worldsForAdmin = append(worldsForAdmin, &dto.WorldForAdminResponse{
			WorldID: v.ID,
			Name:    v.Name,
		})
	}

	return worldsForAdmin, nil
}

func (s mapConfigurationService) ListWorldWithMap(ctx context.Context) ([]*dto.WorldDTO, error) {
	worlds := make([]*entity.World, 0)

	result := s.db.WithContext(ctx).
		Order("id ASC").
		Preload("MapConfigurations", func(db *gorm.DB) *gorm.DB {
			return db.Order("id ASC")
		}).
		Preload("MapConfigurations.Rules").
		Find(&worlds)
	if err := result.Error; err != nil {
		return nil, err
	}

	worldMapper := mapper.NewWorldMapper()
	worldDTOs := make([]*dto.WorldDTO, 0, len(worlds))
	for _, v := range worlds {
		w := worldMapper.ToWorldDTO(v)
		worldDTOs = append(worldDTOs, w)
	}

	return worldDTOs, nil
}

func (s mapConfigurationService) CreateWorld(ctx context.Context, name string) error {
	result := s.db.WithContext(ctx).
		Create(&entity.World{
			Name: name,
		})

	return result.Error
}

func (s mapConfigurationService) UpdateWorld(ctx context.Context, worldID int64, name string) error {
	result := s.db.WithContext(ctx).
		Model(&entity.World{}).
		Where(&entity.World{
			ID: worldID,
		}).
		Update("name", name)

	return result.Error
}

func (s mapConfigurationService) SetMapActive(ctx context.Context, mapID int64, active bool) error {
	result := s.db.WithContext(ctx).
		Model(&entity.MapConfiguration{}).
		Where(&entity.MapConfiguration{
			ID: mapID,
		}).
		Update("active", active)

	return result.Error
}
