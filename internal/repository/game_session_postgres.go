package repository

import (
	"context"
	"database/sql"
	"llg_backend/internal/entity"
	"llg_backend/internal/entity/sqlc_generated"
)

type gameSessionRepository struct {
	sqlc_generated.Querier
}

func NewGameSessionRepository(querier sqlc_generated.Querier) GameSessionRepository {
	return &gameSessionRepository{
		Querier: querier,
	}
}

func (r gameSessionRepository) CreateGameSession(ctx context.Context, arg CreateGameSessionParams) (*entity.PlayerGameSession, error) {
	endDatetimeArg := sql.NullTime{
		Time:  arg.EndDatetime,
		Valid: true,
	}

	if arg.EndDatetime.IsZero() {
		endDatetimeArg.Valid = false
	}

	newCreatedArg := sqlc_generated.CreateGameSessionParams{
		PlayerID:           arg.PlayerID,
		MapConfigurationID: arg.MapConfigurationID,
		StartDatetime:      arg.StartDatetime,
		EndDatetime:        endDatetimeArg,
	}

	gameSessionRow, err := r.Querier.CreateGameSession(ctx, newCreatedArg)
	if err != nil {
		return nil, err
	}

	gameSession := &entity.PlayerGameSession{
		GameSesssionID:     gameSessionRow.ID,
		PlayerID:           gameSessionRow.PlayerID,
		MapConfigurationID: gameSessionRow.MapConfigurationID,
		StartDatetime:      gameSessionRow.StartDatetime,
		EndDatetime:        gameSessionRow.EndDatetime.Time,
		GameHistory:        make([]*entity.PlayHistory, 0),
	}

	return gameSession, nil
}
