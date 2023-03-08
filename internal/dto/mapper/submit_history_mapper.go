package mapper

import (
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
)

type SubmitHistoryMapper struct{}

func NewSubmitHistoryMapper() SubmitHistoryMapper {
	return SubmitHistoryMapper{}
}

func (m SubmitHistoryMapper) ToDTO(submitHistory *entity.SubmitHistory) *dto.SubmitHistoryResponse {
	stateValueMapper := NewStateValueMapper()
	submitHistoryRuleMapper := NewSubmitHistoryRuleMapper()
	commandNodeMapper := NewCommandNodeMapper()
	commandEdgeMapper := NewCommandEdgeMapper()

	submitHistoryRuleDTOs := make([]*dto.SubmitHistoryRuleResponse, 0, len(submitHistory.SubmitHistoryRules))
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

	submitHistoryDTO := &dto.SubmitHistoryResponse{
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

func (m SubmitHistoryMapper) ToEntity(submitHistoryDTO *dto.SubmitHistoryRequest) *entity.SubmitHistory {
	stateValueMapper := NewStateValueMapper()
	submitHistoryRuleMapper := NewSubmitHistoryRuleMapper()
	commandNodeMapper := NewCommandNodeMapper()
	commandEdgeMapper := NewCommandEdgeMapper()

	submitHistoryRules := make([]*entity.SubmitHistoryRule, 0, len(submitHistoryDTO.SubmitHistoryRules))
	for _, v := range submitHistoryDTO.SubmitHistoryRules {
		submitHistoryRule := submitHistoryRuleMapper.ToEntity(v)
		submitHistoryRules = append(submitHistoryRules, submitHistoryRule)
	}

	commandNodes := make([]*entity.CommandNode, 0, len(submitHistoryDTO.CommandNodes))
	for _, v := range submitHistoryDTO.CommandNodes {
		commandNode := commandNodeMapper.ToEntity(v)
		commandNodes = append(commandNodes, commandNode)
	}

	commandEdges := make([]*entity.CommandEdge, 0, len(submitHistoryDTO.CommandEdges))
	for _, v := range submitHistoryDTO.CommandEdges {
		commandEdge := commandEdgeMapper.ToEntity(v)
		commandEdges = append(commandEdges, commandEdge)
	}

	submitHistory := &entity.SubmitHistory{
		IsFinited:          submitHistoryDTO.IsFinited,
		IsCompleted:        submitHistoryDTO.IsCompleted,
		CommandMedal:       submitHistoryDTO.CommandMedal,
		ActionMedal:        submitHistoryDTO.ActionMedal,
		SubmitDatetime:     submitHistoryDTO.SubmitDatetime,
		StateValue:         stateValueMapper.ToEntity(submitHistoryDTO.StateValue),
		SubmitHistoryRules: submitHistoryRules,
		CommandNodes:       commandNodes,
		CommandEdges:       commandEdges,
	}

	return submitHistory
}
