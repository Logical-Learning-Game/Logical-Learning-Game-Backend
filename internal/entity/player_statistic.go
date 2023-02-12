package entity

import (
	"context"
	"llg_backend/internal/entity/sqlc_generated"
	"time"
)

type GameHistoryParams struct {
	ActionStep      int                      `json:"action_step"`
	NumberOfCommand int                      `json:"number_of_command"`
	IsFinited       bool                     `json:"is_finited"`
	IsCompleted     bool                     `json:"is_completed"`
	CommandMedal    sqlc_generated.MedalType `json:"command_medal"`
	ActionMedal     sqlc_generated.MedalType `json:"action_medal"`
	SubmitDatetime  time.Time                `json:"submit_datetime"`
	StateValue      *StateValueParams        `json:"state_value"`
	Rules           []*RuleParams            `json:"rules"`
}

type StateValueParams struct {
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

type RuleParams struct {
	MapConfigRuleID int64 `json:"map_configuration_rule_id"`
	IsPass          bool  `json:"is_pass"`
}

type CreateSessionHistoryParams struct {
	PlayerID           string               `json:"player_id"`
	MapConfigurationID int64                `json:"map_configuration_id"`
	StartDatetime      time.Time            `json:"start_datetime"`
	EndDatetime        time.Time            `json:"end_datetime"`
	GameHistories      []*GameHistoryParams `json:"game_histories"`
}

type PlayerStatisticService interface {
	CreateSessionHistory(ctx context.Context, arg CreateSessionHistoryParams) (*GameSession, error)
}
