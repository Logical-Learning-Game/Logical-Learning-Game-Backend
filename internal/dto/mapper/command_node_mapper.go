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
		NodeIndex: int(commandNode.Index),
		Type:      commandNode.Type,
		InGamePosition: dto.Vector2FloatDTO{
			X: commandNode.InGamePosition.X,
			Y: commandNode.InGamePosition.Y,
		},
	}

	return commandNodeDTO
}
