package repository

import (
	"context"
	"llg_backend/internal/entity"
	"llg_backend/internal/entity/sqlc_generated"
)

type playHistoryRepository struct {
	sqlc_generated.Querier
}

func NewPlayHistory(querier sqlc_generated.Querier) PlayHistoryRepository {
	return &playHistoryRepository{
		Querier: querier,
	}
}

func (r playHistoryRepository) CreatePlayHistory(ctx context.Context, arg CreatePlayHistoryParams) (*entity.PlayHistory, error) {
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
