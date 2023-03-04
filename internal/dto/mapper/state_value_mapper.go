package mapper

import (
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
)

type StateValueMapper struct{}

func NewStateValueMapper() *StateValueMapper {
	return &StateValueMapper{}
}

func (m StateValueMapper) ToDTO(stateValue *entity.StateValue) *dto.StateValueDTO {
	stateValueDTO := &dto.StateValueDTO{
		CommandCount:          int(stateValue.CommandCount),
		ForwardCommandCount:   int(stateValue.ForwardCommandCount),
		RightCommandCount:     int(stateValue.RightCommandCount),
		BackCommandCount:      int(stateValue.BackCommandCount),
		LeftCommandCount:      int(stateValue.LeftCommandCount),
		ConditionCommandCount: int(stateValue.ConditionCommandCount),
		ActionCount:           int(stateValue.ActionCount),
		ForwardActionCount:    int(stateValue.ForwardActionCount),
		RightActionCount:      int(stateValue.RightActionCount),
		BackActionCount:       int(stateValue.BackActionCount),
		LeftActionCount:       int(stateValue.LeftActionCount),
		ConditionActionCount:  int(stateValue.ConditionActionCount),
	}

	return stateValueDTO
}
