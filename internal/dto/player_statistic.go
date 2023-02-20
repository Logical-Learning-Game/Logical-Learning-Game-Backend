package dto

import (
	"llg_backend/internal/entity"
	"llg_backend/internal/entity/nullable"
	"time"
)

type WorldDTO struct {
	WorldID   int64                  `json:"world_id"`
	WorldName string                 `json:"world_name"`
	Maps      []*MapConfigurationDTO `json:"maps"`
}

type MapConfigurationDTO struct {
	MapID                      int64               `json:"map_id"`
	MapName                    string              `json:"map_name"`
	Tile                       [][]int             `json:"tile"`
	MapImagePath               nullable.NullString `json:"map_image_path"`
	Difficulty                 entity.Difficulty   `json:"difficulty"`
	StarRequirement            int                 `json:"star_requirement"`
	LeastSolvableCommandGold   int                 `json:"least_solvable_command_gold"`
	LeastSolvableCommandSilver int                 `json:"least_solvable_command_silver"`
	LeastSolvableCommandBronze int                 `json:"least_solvable_command_bronze"`
	LeastSolvableActionGold    int                 `json:"least_solvable_action_gold"`
	LeastSolvableActionSilver  int                 `json:"least_solvable_action_silver"`
	LeastSolvableActionBronze  int                 `json:"least_solvable_action_bronze"`
	Rules                      []*RuleDTO          `json:"rules"`
	IsPass                     bool                `json:"is_pass"`
	TopHistory                 *SubmitHistoryDTO   `json:"top_history"`
}

type RuleDTO struct {
	MapRuleID  int64            `json:"map_rule_id"`
	RuleName   string           `json:"rule_name"`
	Theme      entity.RuleTheme `json:"rule_theme"`
	Parameters []int            `json:"parameters"`
}

type SubmitHistoryDTO struct {
	ActionStep         int                     `json:"action_step"`
	NumberOfCommand    int                     `json:"number_of_command"`
	IsFinited          bool                    `json:"is_finited"`
	IsCompleted        bool                    `json:"is_completed"`
	CommandMedal       entity.MedalType        `json:"command_medal"`
	ActionMedal        entity.MedalType        `json:"action_medal"`
	SubmitDatetime     time.Time               `json:"submit_datetime"`
	StateValue         *StateValueDTO          `json:"state_value"`
	SubmitHistoryRules []*SubmitHistoryRuleDTO `json:"rules"`
	CommandNodes       []*CommandNodeDTO       `json:"command_nodes"`
	CommandEdges       []*CommandEdgeDTO       `json:"command_edges"`
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

type SubmitHistoryRuleDTO struct {
	MapRuleID int64 `json:"map_rule_id"`
	IsPass    bool  `json:"is_pass"`
}

type CommandNodeDTO struct {
	NodeIndex      int                    `json:"node_index"`
	Type           entity.CommandNodeType `json:"type"`
	InGamePosition Vector2FloatDTO        `json:"in_game_position"`
}

type CommandEdgeDTO struct {
	SourceNodeIndex      int                    `json:"source_node_index"`
	DestinationNodeIndex int                    `json:"destination_node_index"`
	Type                 entity.CommandEdgeType `json:"type"`
}

type Vector2FloatDTO struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type CreateGameSessionRequestDTO struct {
	MapConfigurationID int64               `json:"map_id"`
	StartDatetime      time.Time           `json:"start_datetime"`
	EndDatetime        nullable.NullTime   `json:"end_datetime"`
	SubmitHistories    []*SubmitHistoryDTO `json:"submit_histories"`
}
