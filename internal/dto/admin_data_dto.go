package dto

import (
	"llg_backend/internal/entity"
	"llg_backend/internal/entity/nullable"
	"time"
)

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

type WorldForAdminResponse struct {
	WorldID int64  `json:"world_id"`
	Name    string `json:"world_name"`
}

type UpdateWorldRequest struct {
	Name string `json:"world_name"`
}

type CreateWorldRequest struct {
	Name string `json:"world_name"`
}

type SetMapActiveRequest struct {
	Active bool `json:"active"`
}

type CreateMapRequest struct {
	WorldID                    int64                `json:"world_id"`
	MapName                    string               `json:"map_name"`
	Tile                       []int                `json:"tile"`
	Height                     int                  `json:"height"`
	Width                      int                  `json:"width"`
	StartPlayerDirection       entity.MapDirection  `json:"start_player_direction"`
	StartPlayerPositionX       int                  `json:"start_player_position_x"`
	StartPlayerPositionY       int                  `json:"start_player_position_y"`
	GoalPositionX              int                  `json:"goal_position_x"`
	GoalPositionY              int                  `json:"goal_position_y"`
	Difficulty                 entity.Difficulty    `json:"difficulty"`
	StarRequirement            int                  `json:"star_requirement"`
	LeastSolvableCommandGold   int                  `json:"least_solvable_command_gold"`
	LeastSolvableCommandSilver int                  `json:"least_solvable_command_silver"`
	LeastSolvableCommandBronze int                  `json:"least_solvable_command_bronze"`
	LeastSolvableActionGold    int                  `json:"least_solvable_action_gold"`
	LeastSolvableActionSilver  int                  `json:"least_solvable_action_silver"`
	LeastSolvableActionBronze  int                  `json:"least_solvable_action_bronze"`
	Rules                      []*CreateRuleRequest `json:"rules"`
}

type CreateRuleRequest struct {
	RuleName   string           `json:"rule_name"`
	Theme      entity.RuleTheme `json:"rule_theme"`
	Parameters []int            `json:"parameters"`
}

type WorldWithMapForAdminResponse struct {
	WorldID   int64                          `json:"world_id"`
	WorldName string                         `json:"world_name"`
	Maps      []*MapConfigurationForAdminDTO `json:"maps"`
}

type MapConfigurationForAdminDTO struct {
	MapID                      int64               `json:"map_id"`
	WorldID                    int64               `json:"world_id"`
	MapName                    string              `json:"map_name"`
	Tile                       []int               `json:"tile"`
	Height                     int                 `json:"height"`
	Width                      int                 `json:"width"`
	StartPlayerDirection       entity.MapDirection `json:"start_player_direction"`
	StartPlayerPositionX       int                 `json:"start_player_position_x"`
	StartPlayerPositionY       int                 `json:"start_player_position_y"`
	GoalPositionX              int                 `json:"goal_position_x"`
	GoalPositionY              int                 `json:"goal_position_y"`
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
	Active                     bool                `json:"active"`
}
