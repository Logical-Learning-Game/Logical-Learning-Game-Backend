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
	mapConfigDTOs := make([]*dto.MapConfigurationDTO, 0, len(world.MapConfigurationForPlayers))
	for _, mapConfig := range world.MapConfigurationForPlayers {
		pqInt32Tile := mapConfig.MapConfiguration.Tile

		intTile := make([]int, len(pqInt32Tile))
		for i := range pqInt32Tile {
			intTile[i] = int(pqInt32Tile[i])
		}

		mapConfiguration := mapConfig.MapConfiguration

		rules := make([]*dto.RuleDTO, 0, len(mapConfiguration.Rules))
		for _, rule := range mapConfiguration.Rules {
			pqInt32Parameter := rule.Parameters
			intParameter := make([]int, len(pqInt32Parameter))
			for i := range pqInt32Parameter {
				intParameter[i] = int(pqInt32Parameter[i])
			}

			ruleDTO := &dto.RuleDTO{
				MapRuleID:  rule.ID,
				RuleName:   rule.RuleName,
				Theme:      rule.Theme,
				Parameters: intParameter,
			}

			rules = append(rules, ruleDTO)
		}

		mapConfigDTO := &dto.MapConfigurationDTO{
			MapID:                      mapConfiguration.ID,
			MapName:                    mapConfiguration.ConfigName,
			Tile:                       intTile,
			Height:                     int(mapConfig.MapConfiguration.Height),
			Width:                      int(mapConfig.MapConfiguration.Width),
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

		mapConfigDTOs = append(mapConfigDTOs, mapConfigDTO)
	}

	worldDTO := &dto.WorldDTO{
		WorldID:   world.ID,
		WorldName: world.Name,
		Maps:      mapConfigDTOs,
	}

	return worldDTO
}
