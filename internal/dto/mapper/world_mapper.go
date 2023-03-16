package mapper

import (
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
	"llg_backend/internal/entity/nullable"
)

type WorldMapper struct{}

func NewWorldMapper() WorldMapper {
	return WorldMapper{}
}

func (m WorldMapper) ToDTO(world *entity.World) *dto.WorldDTO {
	if world == nil {
		return nil
	}

	mapConfigDTOs := make([]*dto.MapConfigurationDTO, 0, len(world.MapConfigurations))
	ruleMapper := NewRuleMapper()
	for _, mapConfig := range world.MapConfigurations {
		pqInt32Tile := mapConfig.Tile

		intTile := make([]int, len(pqInt32Tile))
		for i := range pqInt32Tile {
			intTile[i] = int(pqInt32Tile[i])
		}

		rules := make([]*dto.RuleDTO, 0, len(mapConfig.Rules))
		for _, rule := range mapConfig.Rules {
			r := ruleMapper.ToDTO(rule)
			rules = append(rules, r)
		}

		mapConfigDTO := &dto.MapConfigurationDTO{
			MapID:                      mapConfig.ID,
			MapName:                    mapConfig.ConfigName,
			Tile:                       intTile,
			Height:                     int(mapConfig.Height),
			Width:                      int(mapConfig.Width),
			MapImagePath:               nullable.NullString{NullString: mapConfig.MapImagePath},
			Difficulty:                 mapConfig.Difficulty,
			StarRequirement:            int(mapConfig.StarRequirement),
			LeastSolvableCommandGold:   int(mapConfig.LeastSolvableCommandGold),
			LeastSolvableCommandSilver: int(mapConfig.LeastSolvableCommandSilver),
			LeastSolvableCommandBronze: int(mapConfig.LeastSolvableCommandBronze),
			LeastSolvableActionGold:    int(mapConfig.LeastSolvableActionGold),
			LeastSolvableActionSilver:  int(mapConfig.LeastSolvableActionSilver),
			LeastSolvableActionBronze:  int(mapConfig.LeastSolvableActionBronze),
			Rules:                      rules,
		}

		mapConfigDTOs = append(mapConfigDTOs, mapConfigDTO)
	}

	worldDTO := &dto.WorldDTO{
		WorldID:   world.ID,
		WorldName: world.Name,
		Maps:      mapConfigDTOs,
	}

	return worldDTO
}
