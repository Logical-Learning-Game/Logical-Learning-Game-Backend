package mapper

import (
	"github.com/lib/pq"
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
	"llg_backend/internal/entity/nullable"
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

	intTile := pqInt32ArrayToIntSlice(mapConfiguration.Tile)

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

func pqInt32ArrayToIntSlice(pqInt32Array pq.Int32Array) []int {
	intSlice := make([]int, len(pqInt32Array))

	for i := range pqInt32Array {
		intSlice[i] = int(pqInt32Array[i])
	}

	return intSlice
}
