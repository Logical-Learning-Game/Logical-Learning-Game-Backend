package mapper

import (
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
)

type SubmitHistoryMapper struct{}

func NewSubmitHistoryMapper() SubmitHistoryMapper {
	return SubmitHistoryMapper{}
}

func (m SubmitHistoryMapper) ToDTO(submitHistory *entity.SubmitHistory) *dto.SubmitHistoryDTO {
	stateValueMapper := NewStateValueMapper()
	submitHistoryRuleMapper := NewSubmitHistoryRuleMapper()
	commandNodeMapper := NewCommandNodeMapper()
	commandEdgeMapper := NewCommandEdgeMapper()

	submitHistoryRuleDTOs := make([]*dto.SubmitHistoryRuleDTO, 0, len(submitHistory.SubmitHistoryRules))
	for _, v := range submitHistory.SubmitHistoryRules {
		submitHistoryRuleDTO := submitHistoryRuleMapper.ToDTO(v)
		submitHistoryRuleDTOs = append(submitHistoryRuleDTOs, submitHistoryRuleDTO)
	}

	commandNodeDTOs := make([]*dto.CommandNodeDTO, 0, len(submitHistory.CommandNodes))
	for _, v := range submitHistory.CommandNodes {
		commandNodeDTO := commandNodeMapper.ToDTO(v)
		commandNodeDTOs = append(commandNodeDTOs, commandNodeDTO)
	}

	commandEdgeDTOs := make([]*dto.CommandEdgeDTO, 0, len(submitHistory.CommandEdges))
	for _, v := range submitHistory.CommandEdges {
		commandEdgeDTO := commandEdgeMapper.ToDTO(v)
		commandEdgeDTOs = append(commandEdgeDTOs, commandEdgeDTO)
	}

	submitHistoryDTO := &dto.SubmitHistoryDTO{
		IsFinited:          submitHistory.IsFinited,
		IsCompleted:        submitHistory.IsCompleted,
		CommandMedal:       submitHistory.CommandMedal,
		ActionMedal:        submitHistory.ActionMedal,
		SubmitDatetime:     submitHistory.SubmitDatetime,
		StateValue:         stateValueMapper.ToDTO(submitHistory.StateValue),
		SubmitHistoryRules: submitHistoryRuleDTOs,
		CommandNodes:       commandNodeDTOs,
		CommandEdges:       commandEdgeDTOs,
	}

	return submitHistoryDTO
}
