package mapper

import (
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
)

type SubmitHistoryRuleMapper struct{}

func NewSubmitHistoryRuleMapper() SubmitHistoryRuleMapper {
	return SubmitHistoryRuleMapper{}
}

func (m SubmitHistoryRuleMapper) ToDTO(submitHistoryRule *entity.SubmitHistoryRule) *dto.SubmitHistoryRuleDTO {
	submitHistoryRuleDTO := &dto.SubmitHistoryRuleDTO{
		MapRuleID: submitHistoryRule.MapConfigurationRuleID,
		IsPass:    submitHistoryRule.IsPass,
	}

	return submitHistoryRuleDTO
}
