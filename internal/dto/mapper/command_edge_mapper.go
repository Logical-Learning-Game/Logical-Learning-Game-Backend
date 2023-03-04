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
	commandEdgeDTO := &dto.CommandEdgeDTO{
		SourceNodeIndex:      int(commandEdge.SourceNodeIndex),
		DestinationNodeIndex: int(commandEdge.DestinationNodeIndex),
		Type:                 commandEdge.Type,
	}

	return commandEdgeDTO
}
