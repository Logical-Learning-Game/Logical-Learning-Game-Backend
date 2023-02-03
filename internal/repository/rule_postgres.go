package repository

import (
	"context"
	"llg_backend/internal/entity"
	"llg_backend/internal/entity/sqlc_generated"
)

type ruleRepository struct {
	sqlc_generated.Querier
}

func NewRuleRepository(querier sqlc_generated.Querier) RuleRepository {
	return &ruleRepository{
		Querier: querier,
	}
}

func (r *ruleRepository) ListFromMapConfigurationIDs(ctx context.Context, mapConfigurationIDs []int64) ([]*entity.MapRule, error) {
	ruleRows, err := r.Querier.GetRuleFromMapConfigurationIDs(ctx, mapConfigurationIDs)
	if err != nil {
		return nil, err
	}

	mapRules := make([]*entity.MapRule, 0)
	for _, row := range ruleRows {

		intSlice := make([]int, 0)
		for _, val := range row.Parameters {
			intSlice = append(intSlice, int(val))
		}

		mapRules = append(mapRules, &entity.MapRule{
			MapConfigID: row.MapConfigurationID,
			Type:        row.Rule,
			Theme:       row.Theme,
			Parameters:  intSlice,
		})
	}

	return mapRules, nil
}
