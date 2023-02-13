package entity

import (
	"context"
	"llg_backend/internal/entity/sqlc_generated"
	"time"
)

type GameSession struct {
	ID                 int64          `json:"game_session_id"`
	PlayerID           string         `json:"player_id"`
	MapConfigurationID int64          `json:"map_configuration_id"`
	StartDatetime      time.Time      `json:"start_datetime"`
	EndDatetime        time.Time      `json:"end_datetime"`
	GameHistory        []*PlayHistory `json:"game_histories"`
}

type GameSessionRepository interface {
	CreateGameSession(ctx context.Context, arg CreateGameSessionParams) (*GameSession, error)
}

type CreateGameSessionParams struct {
	PlayerID           string
	MapConfigurationID int64
	StartDatetime      time.Time
	EndDatetime        time.Time
}

type PlayHistory struct {
	ID              int64                    `json:"play_history_id"`
	GameSessionID   int64                    `json:"game_session_id"`
	ActionStep      int                      `json:"action_step"`
	NumberOfCommand int                      `json:"number_of_command"`
	IsFinited       bool                     `json:"is_finited"`
	IsCompleted     bool                     `json:"is_completed"`
	CommandMedal    sqlc_generated.MedalType `json:"command_medal"`
	ActionMedal     sqlc_generated.MedalType `json:"action_medal"`
	SubmitDatetime  time.Time                `json:"submit_datetime"`
	StateValue      *StateValue              `json:"state_value"`
	Rules           []*RuleHistory           `json:"rules"`
	CommandNodes    []*CommandNode           `json:"command_nodes"`
	CommandEdges    []*CommandEdge           `json:"command_edges"`
}

type PlayHistoryRepository interface {
	//ListFromGameSessionID(ctx context.Context, gameSessionID int64) ([]*entity.PlayHistory, error)
	//ListFromMapConfigurationID(ctx context.Context, mapConfigurationID int64) ([]*entity.PlayHistory, error)
	CreatePlayHistory(ctx context.Context, arg CreatePlayHistoryParams) (*PlayHistory, error)
	CreateRuleHistory(ctx context.Context, arg CreateRuleHistoryParams) (*RuleHistory, error)
	CreateStateValue(ctx context.Context, arg CreateStateValueParams) (*StateValue, error)
	CreateCommandNode(ctx context.Context, arg CreateCommandNodeParams) (*CommandNode, error)
	CreateCommandEdge(ctx context.Context, arg CreateCommandEdgeParams) (*CommandEdge, error)
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

type CreateCommandNodeParams struct {
	PlayHistoryID  int64
	Type           sqlc_generated.CommandNodeType
	InGamePosition Vector2Float
}

type CreateCommandEdgeParams struct {
	SourceNodeID      int64
	DestinationNodeID int64
	Type              sqlc_generated.CommandEdgeType
}

type RuleHistory struct {
	MapConfigRuleID int64 `json:"map_configuration_rule_id"`
	PlayHistoryID   int64 `json:"-"`
	IsPass          bool  `json:"is_pass"`
}

type StateValue struct {
	CommandCount          int `json:"command_count"`
	ForwardCommandCount   int `json:"forward_command_count"`
	RightCommandCount     int `json:"right_command_count"`
	BackCommandCount      int `json:"back_command_count"`
	LeftCommandCount      int `json:"left_command_count"`
	ConditionCommandCount int `json:"condition_command_count"`
	ActionCount           int `json:"action_count"`
	ForwardActionCount    int `json:"forward_action_count"`
	RightActionCount      int `json:"right_action_count"`
	BackActionCount       int `json:"back_action_count"`
	LeftActionCount       int `json:"left_action_count"`
	ConditionActionCount  int `json:"condition_action_count"`
}

type CommandNode struct {
	ID             int64                          `json:"command_node_id"`
	PlayHistoryID  int64                          `json:"-"`
	Type           sqlc_generated.CommandNodeType `json:"type"`
	InGamePosition Vector2Float                   `json:"in_game_position"`
}

type CommandEdge struct {
	SourceNodeID      int64                          `json:"source_node_id"`
	DestinationNodeID int64                          `json:"destination_node_id"`
	Type              sqlc_generated.CommandEdgeType `json:"type"`
}
