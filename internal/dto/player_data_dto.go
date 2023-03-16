package dto

import (
	"llg_backend/internal/entity"
	"llg_backend/internal/entity/nullable"
	"time"
)

type SubmitHistoryRequest struct {
	IsFinited          bool                        `json:"is_finited"`
	IsCompleted        bool                        `json:"is_completed"`
	CommandMedal       entity.MedalType            `json:"command_medal"`
	ActionMedal        entity.MedalType            `json:"action_medal"`
	SubmitDatetime     time.Time                   `json:"submit_datetime"`
	StateValue         *StateValueDTO              `json:"state_value"`
	SubmitHistoryRules []*SubmitHistoryRuleRequest `json:"rules"`
	CommandNodes       []*CommandNodeDTO           `json:"command_nodes"`
	CommandEdges       []*CommandEdgeDTO           `json:"command_edges"`
}

type SubmitHistoryResponse struct {
	IsFinited          bool                         `json:"is_finited"`
	IsCompleted        bool                         `json:"is_completed"`
	CommandMedal       entity.MedalType             `json:"command_medal"`
	ActionMedal        entity.MedalType             `json:"action_medal"`
	SubmitDatetime     time.Time                    `json:"submit_datetime"`
	StateValue         *StateValueDTO               `json:"state_value"`
	SubmitHistoryRules []*SubmitHistoryRuleResponse `json:"rules"`
	CommandNodes       []*CommandNodeDTO            `json:"command_nodes"`
	CommandEdges       []*CommandEdgeDTO            `json:"command_edges"`
}

type StateValueDTO struct {
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

type SubmitHistoryRuleRequest struct {
	MapRuleID int64 `json:"map_rule_id"`
	IsPass    bool  `json:"is_pass"`
}

type SubmitHistoryRuleResponse struct {
	MapRuleID int64            `json:"map_rule_id"`
	Theme     entity.RuleTheme `json:"theme"`
	IsPass    bool             `json:"is_pass"`
}

type CommandNodeDTO struct {
	Index     int                    `json:"index"`
	Type      entity.CommandNodeType `json:"type"`
	PositionX float32                `json:"x"`
	PositionY float32                `json:"y"`
}

type CommandEdgeDTO struct {
	SourceNodeIndex      int                    `json:"source_node_index"`
	DestinationNodeIndex int                    `json:"destination_node_index"`
	Type                 entity.CommandEdgeType `json:"type"`
}

type SessionHistoryRequest struct {
	MapConfigurationID int64                   `json:"map_id"`
	StartDatetime      time.Time               `json:"start_datetime"`
	EndDatetime        nullable.NullTime       `json:"end_datetime"`
	SubmitHistories    []*SubmitHistoryRequest `json:"submit_histories"`
}

type SessionHistoryForGameResponse struct {
	MapConfigurationID int64                    `json:"map_id"`
	StartDatetime      time.Time                `json:"start_datetime"`
	EndDatetime        nullable.NullTime        `json:"end_datetime"`
	SubmitHistories    []*SubmitHistoryResponse `json:"submit_histories"`
}

type TopSubmitHistoryRequest struct {
	MapConfigurationID int64                 `json:"map_id"`
	SubmitHistory      *SubmitHistoryRequest `json:"top_submit_history"`
}

type TopSubmitHistoryResponse struct {
	MapConfigurationID int64                  `json:"map_id"`
	SubmitHistory      *SubmitHistoryResponse `json:"top_submit_history"`
}

type PlayerDataDTO struct {
	SessionHistories   []*SessionHistoryForGameResponse
	TopSubmitHistories []*TopSubmitHistoryResponse
}

type PlayerDataResponse struct {
	PlayerID           string                              `json:"player_id"`
	SessionHistories   []*SessionHistoryWithStatusResponse `json:"session_histories"`
	TopSubmitHistories map[int64]*SubmitHistoryResponse    `json:"top_submits"`
}

type SessionHistoryWithStatusResponse struct {
	SessionHistory *SessionHistoryForGameResponse `json:"session"`
	Status         bool                           `json:"status"`
}

type LinkAccountRequest struct {
	PlayerID string `json:"player_id"`
	Email    string `json:"email"`
}

type PlayerInfoResponse struct {
	PlayerID string `json:"player_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

type SessionDataForAdminResponse struct {
	SessionID       int64                            `json:"session_id"`
	WorldID         int64                            `json:"world_id"`
	WorldName       string                           `json:"world_name"`
	MapID           int64                            `json:"map_id"`
	MapName         string                           `json:"map_name"`
	StartDatetime   time.Time                        `json:"start_datetime"`
	EndDatetime     nullable.NullTime                `json:"end_datetime"`
	SubmitHistories []*SubmitHistoryForAdminResponse `json:"submit_histories"`
}

type SubmitHistoryForAdminResponse struct {
	SubmitHistoryID    int64                                `json:"submit_history_id"`
	IsFinited          bool                                 `json:"is_finited"`
	IsCompleted        bool                                 `json:"is_completed"`
	CommandMedal       entity.MedalType                     `json:"command_medal"`
	ActionMedal        entity.MedalType                     `json:"action_medal"`
	SubmitDatetime     time.Time                            `json:"submit_datetime"`
	StateValue         *StateValueDTO                       `json:"state_value"`
	SubmitHistoryRules []*SubmitHistoryRuleForAdminResponse `json:"rules"`
	CommandNodes       []*CommandNodeDTO                    `json:"command_nodes"`
	CommandEdges       []*CommandEdgeDTO                    `json:"command_edges"`
}

type SubmitHistoryRuleForAdminResponse struct {
	Rule   *RuleDTO `json:"rule"`
	IsPass bool     `json:"is_pass"`
}

type MapOfPlayerInfoForAdminResponse struct {
	MapForPlayerID   int64                          `json:"map_for_player_id"`
	WorldID          int64                          `json:"world_id"`
	WorldName        string                         `json:"world_name"`
	MapID            int64                          `json:"map_id"`
	MapName          string                         `json:"map_name"`
	IsPass           bool                           `json:"is_pass"`
	TopSubmitHistory *SubmitHistoryForAdminResponse `json:"top_submit_history"`
}
