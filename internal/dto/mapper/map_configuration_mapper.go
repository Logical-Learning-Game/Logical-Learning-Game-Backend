package mapper

import (
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
	"llg_backend/internal/entity/nullable"
	"llg_backend/pkg/utility"
)

type MapConfigurationMapper struct{}

func NewMapConfigurationMapper() MapConfigurationMapper {
	return MapConfigurationMapper{}
}

func (m MapConfigurationMapper) ToMapConfigurationDTO(mapConfiguration *entity.MapConfiguration) *dto.MapConfigurationDTO {
	if mapConfiguration == nil {
		return nil
	}

	ruleMapper := NewRuleMapper()

	intTile := utility.PqInt32ArrayToIntSlice(mapConfiguration.Tile)

	rules := make([]*dto.RuleDTO, 0, len(mapConfiguration.Rules))
	for _, rule := range mapConfiguration.Rules {
		r := ruleMapper.ToDTO(rule)
		rules = append(rules, r)
	}

	mapConfigDTO := &dto.MapConfigurationDTO{
		MapID:                      mapConfiguration.ID,
		MapName:                    mapConfiguration.ConfigName,
		Tile:                       intTile,
		Height:                     int(mapConfiguration.Height),
		Width:                      int(mapConfiguration.Width),
		MapImagePath:               nullable.NullString{NullString: mapConfiguration.MapImagePath},
		Difficulty:                 mapConfiguration.Difficulty,
		StarRequirement:            int(mapConfiguration.StarRequirement),
		LeastSolvableCommandGold:   int(mapConfiguration.LeastSolvableCommandGold),
		LeastSolvableCommandSilver: int(mapConfiguration.LeastSolvableCommandSilver),
		LeastSolvableCommandBronze: int(mapConfiguration.LeastSolvableCommandBronze),
		LeastSolvableActionGold:    int(mapConfiguration.LeastSolvableActionGold),
		LeastSolvableActionSilver:  int(mapConfiguration.LeastSolvableActionSilver),
		LeastSolvableActionBronze:  int(mapConfiguration.LeastSolvableActionBronze),
		Rules:                      rules,
	}

	return mapConfigDTO
}

func (m MapConfigurationMapper) ToMapConfigurationForAdminDTO(mapConfiguration *entity.MapConfiguration) *dto.MapConfigurationForAdminDTO {
	if mapConfiguration == nil {
		return nil
	}

	ruleMapper := NewRuleMapper()

	intTile := utility.PqInt32ArrayToIntSlice(mapConfiguration.Tile)

	rules := make([]*dto.RuleDTO, 0, len(mapConfiguration.Rules))
	for _, rule := range mapConfiguration.Rules {
		r := ruleMapper.ToDTO(rule)
		rules = append(rules, r)
	}

	mapConfigForAdmin := &dto.MapConfigurationForAdminDTO{
		MapID:                      mapConfiguration.ID,
		WorldID:                    mapConfiguration.WorldID,
		MapName:                    mapConfiguration.ConfigName,
		Tile:                       intTile,
		Height:                     int(mapConfiguration.Height),
		Width:                      int(mapConfiguration.Width),
		StartPlayerDirection:       mapConfiguration.StartPlayerDirection,
		StartPlayerPositionX:       int(mapConfiguration.StartPlayerPosition.X),
		StartPlayerPositionY:       int(mapConfiguration.StartPlayerPosition.Y),
		GoalPositionX:              int(mapConfiguration.GoalPosition.X),
		GoalPositionY:              int(mapConfiguration.GoalPosition.Y),
		MapImagePath:               nullable.NullString{NullString: mapConfiguration.MapImagePath},
		Difficulty:                 mapConfiguration.Difficulty,
		StarRequirement:            int(mapConfiguration.StarRequirement),
		LeastSolvableCommandGold:   int(mapConfiguration.LeastSolvableCommandGold),
		LeastSolvableCommandSilver: int(mapConfiguration.LeastSolvableCommandSilver),
		LeastSolvableCommandBronze: int(mapConfiguration.LeastSolvableCommandBronze),
		LeastSolvableActionGold:    int(mapConfiguration.LeastSolvableActionGold),
		LeastSolvableActionSilver:  int(mapConfiguration.LeastSolvableActionSilver),
		LeastSolvableActionBronze:  int(mapConfiguration.LeastSolvableActionBronze),
		Rules:                      rules,
		Active:                     mapConfiguration.Active,
	}

	return mapConfigForAdmin
}
