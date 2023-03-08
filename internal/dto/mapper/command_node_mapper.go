package mapper

import (
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
)

type CommandNodeMapper struct{}

func NewCommandNodeMapper() CommandNodeMapper {
	return CommandNodeMapper{}
}

func (m CommandNodeMapper) ToDTO(commandNode *entity.CommandNode) *dto.CommandNodeDTO {
	commandNodeDTO := &dto.CommandNodeDTO{
		Index:     int(commandNode.Index),
		Type:      commandNode.Type,
		PositionX: commandNode.InGamePosition.X,
		PositionY: commandNode.InGamePosition.Y,
	}

	return commandNodeDTO
}

func (m CommandNodeMapper) ToEntity(commandNodeDTO *dto.CommandNodeDTO) *entity.CommandNode {
	commandNode := &entity.CommandNode{
		Index: int32(commandNodeDTO.Index),
		Type:  commandNodeDTO.Type,
		InGamePosition: entity.Vector2Float32{
			X: commandNodeDTO.PositionX,
			Y: commandNodeDTO.PositionY,
		},
	}

	return commandNode
}
