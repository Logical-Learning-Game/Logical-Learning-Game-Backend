package repository

import (
	"context"
	"llg_backend/internal/entity"
	"llg_backend/internal/entity/sqlc_generated"
)

type playHistoryRepository struct {
	sqlc_generated.Querier
}

func NewPlayHistoryRepository(querier sqlc_generated.Querier) entity.PlayHistoryRepository {
	return &playHistoryRepository{
		Querier: querier,
	}
}

func (r playHistoryRepository) CreatePlayHistory(ctx context.Context, arg entity.CreatePlayHistoryParams) (*entity.PlayHistory, error) {
	newCreatedArg := sqlc_generated.CreatePlayHistoryParams{
		GameSessionID:   arg.GameSessionID,
		ActionStep:      int32(arg.ActionStep),
		NumberOfCommand: int32(arg.NumberOfCommand),
		IsFinited:       arg.IsFinited,
		IsCompleted:     arg.IsCompleted,
		CommandMedal:    arg.CommandMedal,
		ActionMedal:     arg.ActionMedal,
		SubmitDatetime:  arg.SubmitDatetime,
	}

	playHistoryRow, err := r.Querier.CreatePlayHistory(ctx, newCreatedArg)
	if err != nil {
		return nil, err
	}

	playHistory := &entity.PlayHistory{
		ID:              playHistoryRow.ID,
		GameSessionID:   playHistoryRow.GameSessionID,
		ActionStep:      int(playHistoryRow.ActionStep),
		NumberOfCommand: int(playHistoryRow.NumberOfCommand),
		IsFinited:       playHistoryRow.IsFinited,
		IsCompleted:     playHistoryRow.IsCompleted,
		ActionMedal:     playHistoryRow.ActionMedal,
		CommandMedal:    playHistoryRow.CommandMedal,
		SubmitDatetime:  playHistoryRow.SubmitDatetime,
		Rules:           make([]*entity.RuleHistory, 0),
	}

	return playHistory, nil
}

func (r playHistoryRepository) CreateRuleHistory(ctx context.Context, arg entity.CreateRuleHistoryParams) (*entity.RuleHistory, error) {
	newCreatedArg := sqlc_generated.CreateRuleHistoryParams{
		PlayHistoryID:          arg.PlayHistoryID,
		MapConfigurationRuleID: arg.MapConfigRuleID,
		IsPass:                 arg.IsPass,
	}

	ruleHistoryRow, err := r.Querier.CreateRuleHistory(ctx, newCreatedArg)
	if err != nil {
		return nil, err
	}

	ruleHistory := &entity.RuleHistory{
		MapConfigRuleID: ruleHistoryRow.MapConfigurationRuleID,
		PlayHistoryID:   ruleHistoryRow.PlayHistoryID,
		Rule:            nil,
		IsPass:          ruleHistoryRow.IsPass,
	}

	return ruleHistory, nil
}

func (r playHistoryRepository) CreateStateValue(ctx context.Context, arg entity.CreateStateValueParams) (*entity.StateValue, error) {
	newCreatedArg := sqlc_generated.CreateStateValueParams{
		PlayHistoryID:         arg.PlayHistoryID,
		CommandCount:          int32(arg.CommandCount),
		ForwardCommandCount:   int32(arg.ForwardCommandCount),
		RightCommandCount:     int32(arg.RightCommandCount),
		BackCommandCount:      int32(arg.BackCommandCount),
		LeftCommandCount:      int32(arg.LeftCommandCount),
		ConditionCommandCount: int32(arg.ConditionCommandCount),
		ActionCount:           int32(arg.ActionCount),
		ForwardActionCount:    int32(arg.ForwardActionCount),
		RightActionCount:      int32(arg.RightActionCount),
		BackActionCount:       int32(arg.BackActionCount),
		LeftActionCount:       int32(arg.LeftActionCount),
		ConditionActionCount:  int32(arg.ConditionActionCount),
	}

	stateValueRow, err := r.Querier.CreateStateValue(ctx, newCreatedArg)
	if err != nil {
		return nil, err
	}

	stateValue := &entity.StateValue{
		CommandCount:          int(stateValueRow.CommandCount),
		ForwardCommandCount:   int(stateValueRow.ForwardCommandCount),
		RightCommandCount:     int(stateValueRow.RightCommandCount),
		BackCommandCount:      int(stateValueRow.BackCommandCount),
		LeftCommandCount:      int(stateValueRow.LeftCommandCount),
		ConditionCommandCount: int(stateValueRow.ConditionCommandCount),
		ActionCount:           int(stateValueRow.ActionCount),
		ForwardActionCount:    int(stateValueRow.ForwardActionCount),
		RightActionCount:      int(stateValueRow.RightActionCount),
		BackActionCount:       int(stateValueRow.BackActionCount),
		LeftActionCount:       int(stateValueRow.LeftActionCount),
		ConditionActionCount:  int(stateValueRow.ConditionActionCount),
	}

	return stateValue, nil
}
