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
	CommandMedal    sqlc_generated.MedalType
	ActionMedal     sqlc_generated.MedalType
	SubmitDatetime  time.Time
}

type CreateRuleHistoryParams struct {
	PlayHistoryID   int64
	MapConfigRuleID int64
	IsPass          bool
}

type CreateStateValueParams struct {
	PlayHistoryID         int64
	CommandCount          int
	ForwardCommandCount   int
	RightCommandCount     int
	BackCommandCount      int
	LeftCommandCount      int
	ConditionCommandCount int
	ActionCount           int
	ForwardActionCount    int
	RightActionCount      int
	BackActionCount       int
	LeftActionCount       int
	ConditionActionCount  int
}

type PlayHistoryRepository interface {
	CreatePlayHistory(ctx context.Context, arg CreatePlayHistoryParams) (*entity.PlayHistory, error)
	CreateRuleHistory(ctx context.Context, arg CreateRuleHistoryParams) (*entity.RuleHistory, error)
	CreateStateValue(ctx context.Context, arg CreateStateValueParams) (*entity.StateValue, error)
}

type UnitOfWork interface {
	Do(ctx context.Context, fn UnitOfWorkBlock) error
}

type UnitOfWorkBlock func(*UnitOfWorkStore) error
