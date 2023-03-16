package mapper

import (
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
)

type SubmitHistoryRuleMapper struct{}

func NewSubmitHistoryRuleMapper() SubmitHistoryRuleMapper {
	return SubmitHistoryRuleMapper{}
}

func (m SubmitHistoryRuleMapper) ToSubmitHistoryRuleResponse(submitHistoryRule *entity.SubmitHistoryRule) *dto.SubmitHistoryRuleResponse {
	if submitHistoryRule == nil {
		return nil
	}

	submitHistoryRuleDTO := &dto.SubmitHistoryRuleResponse{
		MapRuleID: submitHistoryRule.MapConfigurationRuleID,
		Theme:     submitHistoryRule.MapConfigurationRule.Theme,
		IsPass:    submitHistoryRule.IsPass,
	}

	return submitHistoryRuleDTO
}

func (m SubmitHistoryRuleMapper) ToSubmitHistoryRuleForAdminResponse(submitHistoryRule *entity.SubmitHistoryRule) *dto.SubmitHistoryRuleForAdminResponse {
	if submitHistoryRule == nil {
		return nil
	}

	ruleMapper := NewRuleMapper()

	submitHistoryRuleForAdminResponse := &dto.SubmitHistoryRuleForAdminResponse{
		Rule:   ruleMapper.ToDTO(submitHistoryRule.MapConfigurationRule),
		IsPass: submitHistoryRule.IsPass,
	}

	return submitHistoryRuleForAdminResponse
}

func (m SubmitHistoryRuleMapper) ToEntity(submitHistoryRuleDTO *dto.SubmitHistoryRuleRequest) *entity.SubmitHistoryRule {
	if submitHistoryRuleDTO == nil {
		return nil
	}

	submitHistoryRule := &entity.SubmitHistoryRule{
		MapConfigurationRuleID: submitHistoryRuleDTO.MapRuleID,
		IsPass:                 submitHistoryRuleDTO.IsPass,
	}

	return submitHistoryRule
}
