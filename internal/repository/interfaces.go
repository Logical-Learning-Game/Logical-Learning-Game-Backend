package repository

import (
	"context"
	"llg_backend/internal/entity"
	"llg_backend/internal/entity/sqlc_generated"
	"time"
)

type PlayerRepository interface {
	CreateLoginLog(ctx context.Context, playerID string) error
	CreateOrUpdatePlayer(ctx context.Context, playerID, email, name string) error
}

type MapConfigurationRepository interface {
	ListFromPlayerID(ctx context.Context, playerID string) ([]*entity.PlayerStatInMap, error)
}

type ItemRepository interface {
	ListFromMapConfigurationIDs(ctx context.Context, mapConfigurationIDs []int64) ([]*entity.MapItem, error)
}

type DoorRepository interface {
	ListFromMapConfigurationIDs(ctx context.Context, mapConfigurationIDs []int64) ([]*entity.MapDoor, error)
}

type WorldRepository interface {
	ListFromMapConfigurationIDs(ctx context.Context, mapConfigurationIDs []int64) ([]*entity.World, error)
}

type RuleRepository interface {
	ListFromMapConfigurationIDs(ctx context.Context, mapConfigurationIDs []int64) ([]*entity.MapRule, error)
}

type CreateGameSessionParams struct {
	PlayerID           string
	MapConfigurationID int64
	StartDatetime      time.Time
	EndDatetime        time.Time
}

type GameSessionRepository interface {
	CreateGameSession(ctx context.Context, arg CreateGameSessionParams) (*entity.PlayerGameSession, error)
}

type CreatePlayHistoryParams struct {
	GameSessionID   int64
	ActionStep      int
	NumberOfCommand int
	IsFinited       bool
	IsCompleted     bool
	CommandMedal    sqlc_generated.NullMedalType
	ActionMedal     sqlc_generated.NullMedalType
	SubmitDatetime  time.Time
}

type PlayHistoryRepository interface {
	CreatePlayHistory(ctx context.Context, arg CreatePlayHistoryParams) (*entity.PlayHistory, error)
}
