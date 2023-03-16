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
	if stateValue == nil {
		return nil
	}

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

func (m StateValueMapper) ToEntity(stateValueDTO *dto.StateValueDTO) *entity.StateValue {
	if stateValueDTO == nil {
		return nil
	}

	stateValue := &entity.StateValue{
		CommandCount:          int32(stateValueDTO.CommandCount),
		ForwardCommandCount:   int32(stateValueDTO.ForwardCommandCount),
		RightCommandCount:     int32(stateValueDTO.RightCommandCount),
		BackCommandCount:      int32(stateValueDTO.BackCommandCount),
		LeftCommandCount:      int32(stateValueDTO.LeftCommandCount),
		ConditionCommandCount: int32(stateValueDTO.ConditionCommandCount),
		ActionCount:           int32(stateValueDTO.ActionCount),
		ForwardActionCount:    int32(stateValueDTO.ForwardActionCount),
		RightActionCount:      int32(stateValueDTO.RightActionCount),
		BackActionCount:       int32(stateValueDTO.BackActionCount),
		LeftActionCount:       int32(stateValueDTO.LeftActionCount),
		ConditionActionCount:  int32(stateValueDTO.ConditionActionCount),
	}

	return stateValue
}
