package mapper

import (
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
)

type SubmitHistoryRuleMapper struct{}

func NewSubmitHistoryRuleMapper() SubmitHistoryRuleMapper {
	return SubmitHistoryRuleMapper{}
}

func (m SubmitHistoryRuleMapper) ToDTO(submitHistoryRule *entity.SubmitHistoryRule) *dto.SubmitHistoryRuleResponse {
	submitHistoryRuleDTO := &dto.SubmitHistoryRuleResponse{
		MapRuleID: submitHistoryRule.MapConfigurationRuleID,
		Theme:     submitHistoryRule.MapConfigurationRule.Theme,
		IsPass:    submitHistoryRule.IsPass,
	}

	return submitHistoryRuleDTO
}

func (m SubmitHistoryRuleMapper) ToEntity(submitHistoryRuleDTO *dto.SubmitHistoryRuleRequest) *entity.SubmitHistoryRule {
	submitHistoryRule := &entity.SubmitHistoryRule{
		MapConfigurationRuleID: submitHistoryRuleDTO.MapRuleID,
		IsPass:                 submitHistoryRuleDTO.IsPass,
	}

	return submitHistoryRule
}
