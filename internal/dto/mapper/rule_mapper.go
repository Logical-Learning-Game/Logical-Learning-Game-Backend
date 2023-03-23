package mapper

import (
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
)

type RuleMapper struct{}

func NewRuleMapper() RuleMapper {
	return RuleMapper{}
}

func (m RuleMapper) ToDTO(rule *entity.MapConfigurationRule) *dto.RuleDTO {
	if rule == nil {
		return nil
	}

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

	return ruleDTO
}
