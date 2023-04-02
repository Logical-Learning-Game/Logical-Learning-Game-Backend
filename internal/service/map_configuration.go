package service

import (
	"context"
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"llg_backend/internal/dto"
	"llg_backend/internal/dto/mapper"
	"llg_backend/internal/entity"
	"llg_backend/pkg/utility"
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

func (s mapConfigurationService) ListWorldWithMap(ctx context.Context) ([]*dto.WorldWithMapForAdminResponse, error) {
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
	worldWithMapResponse := make([]*dto.WorldWithMapForAdminResponse, 0, len(worlds))
	for _, v := range worlds {
		w := worldMapper.ToWorldWithMapForAdminResponse(v)
		worldWithMapResponse = append(worldWithMapResponse, w)
	}

	return worldWithMapResponse, nil
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

func (s mapConfigurationService) GetMapByID(ctx context.Context, mapID int64) (*dto.MapConfigurationForAdminDTO, error) {
	var mapConfig entity.MapConfiguration

	result := s.db.WithContext(ctx).
		Where(&entity.MapConfiguration{
			ID: mapID,
		}).
		Preload("Rules").
		First(&mapConfig)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMapNotFound
		}
		return nil, err
	}

	mapConfigDTOMapper := mapper.NewMapConfigurationMapper()
	mapConfigForAdminDTO := mapConfigDTOMapper.ToMapConfigurationForAdminDTO(&mapConfig)
	return mapConfigForAdminDTO, nil
}

func (s mapConfigurationService) CreateMap(ctx context.Context, createMapRequest *dto.CreateMapRequest, imagePath string) error {
	txErr := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		mapConfig := &entity.MapConfiguration{
			WorldID:              createMapRequest.WorldID,
			ConfigName:           createMapRequest.MapName,
			Tile:                 utility.IntSliceToPqInt32Array(createMapRequest.Tile),
			Height:               int32(createMapRequest.Height),
			Width:                int32(createMapRequest.Width),
			StartPlayerDirection: createMapRequest.StartPlayerDirection,
			StartPlayerPosition: entity.Vector2Int{
				X: int32(createMapRequest.StartPlayerPositionX),
				Y: int32(createMapRequest.StartPlayerPositionY),
			},
			GoalPosition: entity.Vector2Int{
				X: int32(createMapRequest.GoalPositionX),
				Y: int32(createMapRequest.GoalPositionY),
			},
			MapImagePath: sql.NullString{
				String: imagePath,
				Valid:  imagePath != "",
			},
			Difficulty:                 createMapRequest.Difficulty,
			StarRequirement:            int32(createMapRequest.StarRequirement),
			LeastSolvableCommandGold:   int32(createMapRequest.LeastSolvableCommandGold),
			LeastSolvableCommandSilver: int32(createMapRequest.LeastSolvableCommandSilver),
			LeastSolvableCommandBronze: int32(createMapRequest.LeastSolvableCommandBronze),
			LeastSolvableActionGold:    int32(createMapRequest.LeastSolvableActionGold),
			LeastSolvableActionSilver:  int32(createMapRequest.LeastSolvableActionSilver),
			LeastSolvableActionBronze:  int32(createMapRequest.LeastSolvableActionBronze),
		}

		result := tx.Create(mapConfig)
		if err := result.Error; err != nil {
			return err
		}

		rules := make([]*entity.MapConfigurationRule, 0, len(createMapRequest.Rules))

		for _, v := range createMapRequest.Rules {
			rules = append(rules, &entity.MapConfigurationRule{
				MapConfigurationID: mapConfig.ID,
				RuleName:           v.RuleName,
				Theme:              v.Theme,
				Parameters:         utility.IntSliceToPqInt32Array(v.Parameters),
			})
		}

		result = tx.Create(&rules)
		if err := result.Error; err != nil {
			return err
		}

		return nil
	})

	return txErr
}

func (s mapConfigurationService) UpdateMap(ctx context.Context, mapID int64, createMapRequest *dto.CreateMapRequest, imagePath string) error {
	txErr := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var oldMap entity.MapConfiguration
		result := tx.Where(&entity.MapConfiguration{
			ID: mapID,
		}).First(&oldMap)
		if err := result.Error; err != nil {
			return err
		}

		// if imagePath is empty string then copy old image url to new map entry
		updateMapImagePath := sql.NullString{
			String: imagePath,
			Valid:  true,
		}
		if imagePath == "" {
			if oldMap.MapImagePath.Valid {
				updateMapImagePath.String = oldMap.MapImagePath.String
			} else {
				updateMapImagePath.Valid = false
			}
		}

		// do create new map
		mapConfig := &entity.MapConfiguration{
			WorldID:              createMapRequest.WorldID,
			ConfigName:           createMapRequest.MapName,
			Tile:                 utility.IntSliceToPqInt32Array(createMapRequest.Tile),
			Height:               int32(createMapRequest.Height),
			Width:                int32(createMapRequest.Width),
			StartPlayerDirection: createMapRequest.StartPlayerDirection,
			StartPlayerPosition: entity.Vector2Int{
				X: int32(createMapRequest.StartPlayerPositionX),
				Y: int32(createMapRequest.StartPlayerPositionY),
			},
			GoalPosition: entity.Vector2Int{
				X: int32(createMapRequest.GoalPositionX),
				Y: int32(createMapRequest.GoalPositionY),
			},
			MapImagePath:               updateMapImagePath,
			Difficulty:                 createMapRequest.Difficulty,
			StarRequirement:            int32(createMapRequest.StarRequirement),
			LeastSolvableCommandGold:   int32(createMapRequest.LeastSolvableCommandGold),
			LeastSolvableCommandSilver: int32(createMapRequest.LeastSolvableCommandSilver),
			LeastSolvableCommandBronze: int32(createMapRequest.LeastSolvableCommandBronze),
			LeastSolvableActionGold:    int32(createMapRequest.LeastSolvableActionGold),
			LeastSolvableActionSilver:  int32(createMapRequest.LeastSolvableActionSilver),
			LeastSolvableActionBronze:  int32(createMapRequest.LeastSolvableActionBronze),
		}

		result = tx.Create(mapConfig)
		if err := result.Error; err != nil {
			return err
		}

		rules := make([]*entity.MapConfigurationRule, 0, len(createMapRequest.Rules))

		for _, v := range createMapRequest.Rules {
			rules = append(rules, &entity.MapConfigurationRule{
				MapConfigurationID: mapConfig.ID,
				RuleName:           v.RuleName,
				Theme:              v.Theme,
				Parameters:         utility.IntSliceToPqInt32Array(v.Parameters),
			})
		}

		result = tx.Create(&rules)
		if err := result.Error; err != nil {
			return err
		}

		// disable old map
		oldMap.Active = false
		result = tx.Save(&oldMap)

		return result.Error
	})

	return txErr
}
