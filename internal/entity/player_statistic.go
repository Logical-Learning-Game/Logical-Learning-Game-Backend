package entity

import (
	"context"
	"llg_backend/internal/entity/sqlc_generated"
	"time"
)

type SubmitHistoryRequest struct {
	ActionStep      int                      `json:"action_step"`
	NumberOfCommand int                      `json:"number_of_command"`
	IsFinited       bool                     `json:"is_finited"`
	IsCompleted     bool                     `json:"is_completed"`
	CommandMedal    sqlc_generated.MedalType `json:"command_medal"`
	ActionMedal     sqlc_generated.MedalType `json:"action_medal"`
	SubmitDatetime  time.Time                `json:"submit_datetime"`
	StateValue      *StateValue              `json:"state_value"`
	Rules           []*RuleHistoryRequest    `json:"rules"`
	CommandNodes    []*CommandNodeRequest    `json:"command_nodes"`
	CommandEdges    []*CommandEdgeRequest    `json:"command_edges"`
}

type CommandNodeRequest struct {
	NodeIndex      int                            `json:"node_index"`
	Type           sqlc_generated.CommandNodeType `json:"type"`
	InGamePosition Vector2Float                   `json:"in_game_position"`
}

type CommandEdgeRequest struct {
	SourceNodeIndex  int                            `json:"source_node_index"`
	DestinationIndex int                            `json:"destination_node_index"`
	Type             sqlc_generated.CommandEdgeType `json:"type"`
}

type RuleHistoryRequest struct {
	MapConfigRuleID int64 `json:"map_configuration_rule_id"`
	IsPass          bool  `json:"is_pass"`
}

type CreateSessionHistoryRequest struct {
	PlayerID           string                  `json:"-"`
	MapConfigurationID int64                   `json:"map_configuration_id"`
	StartDatetime      time.Time               `json:"start_datetime"`
	EndDatetime        time.Time               `json:"end_datetime"`
	GameHistories      []*SubmitHistoryRequest `json:"game_histories"`
}

type PlayerStatisticService interface {
	CreateSessionHistory(ctx context.Context, arg CreateSessionHistoryRequest) (*GameSession, error)
}
