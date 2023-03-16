package mapper

import (
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
)

type CommandEdgeMapper struct{}

func NewCommandEdgeMapper() CommandEdgeMapper {
	return CommandEdgeMapper{}
}

func (m CommandEdgeMapper) ToDTO(commandEdge *entity.CommandEdge) *dto.CommandEdgeDTO {
	if commandEdge == nil {
		return nil
	}

	commandEdgeDTO := &dto.CommandEdgeDTO{
		SourceNodeIndex:      int(commandEdge.SourceNodeIndex),
		DestinationNodeIndex: int(commandEdge.DestinationNodeIndex),
		Type:                 commandEdge.Type,
	}

	return commandEdgeDTO
}

func (m CommandEdgeMapper) ToEntity(commandEdgeDTO *dto.CommandEdgeDTO) *entity.CommandEdge {
	if commandEdgeDTO == nil {
		return nil
	}

	commandEdge := &entity.CommandEdge{
		SourceNodeIndex:      int32(commandEdgeDTO.SourceNodeIndex),
		DestinationNodeIndex: int32(commandEdgeDTO.DestinationNodeIndex),
		Type:                 commandEdgeDTO.Type,
	}

	return commandEdge
}
