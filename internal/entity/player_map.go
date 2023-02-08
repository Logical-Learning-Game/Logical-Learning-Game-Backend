package entity

import (
	"llg_backend/internal/entity/sqlc_generated"
	"time"
)

type Vector2Int struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PlayerData struct {
	PlayerID string
	Email    string
	Worlds   []*World
}

type MapItem struct {
	MapConfigID int64                   `json:"map_config_id"`
	Name        string                  `json:"name"`
	Type        sqlc_generated.ItemType `json:"type"`
	Position    Vector2Int              `json:"position"`
}

type MapDoor struct {
	MapConfigID   int64                       `json:"map_config_id"`
	Name          string                      `json:"name"`
	Type          sqlc_generated.DoorType     `json:"type"`
	Position      Vector2Int                  `json:"position"`
	DoorDirection sqlc_generated.MapDirection `json:"door_direction"`
}

type MapRule struct {
	ID          int64                    `json:"map_config_rule_id"`
	MapConfigID int64                    `json:"map_config_id"`
	Type        string                   `json:"type"`
	Theme       sqlc_generated.RuleTheme `json:"theme"`
	Parameters  []int                    `json:"parameters"`
}

type BadgeRequirement struct {
	LeastSolvableCommandGold   int `json:"least_solvable_command_gold"`
	LeastSolvableCommandSilver int `json:"least_solvable_command_silver"`
	LeastSolvableCommandBronze int `json:"least_solvable_command_bronze"`
	LeastSolvableActionGold    int `json:"least_solvable_action_gold"`
	LeastSolvableActionSilver  int `json:"least_solvable_action_silver"`
	LeastSolvableActionBronze  int `json:"least_solvable_action_bronze"`
}

type MapConfiguration struct {
	BadgeRequirement

	ID                   int64                        `json:"map_config_id"`
	WorldID              int64                        `json:"world_id"`
	ConfigName           string                       `json:"config_name"`
	Map                  [][]int                      `json:"map"`
	Height               int                          `json:"height"`
	Width                int                          `json:"width"`
	StartPlayerDirection sqlc_generated.MapDirection  `json:"start_player_direction"`
	StartPlayerPosition  Vector2Int                   `json:"start_player_position"`
	GoalPosition         Vector2Int                   `json:"goal_position"`
	MapImagePath         string                       `json:"map_image_path"`
	Difficulty           sqlc_generated.MapDifficulty `json:"difficulty"`
	StarRequirement      int                          `json:"star_requirement"`
	Items                []*MapItem                   `json:"items"`
	Doors                []*MapDoor                   `json:"doors"`
	Rules                []*MapRule                   `json:"rules"`
}

type PlayerStatInMap struct {
	MapConfig  *MapConfiguration `json:"map"`
	IsPass     bool              `json:"is_pass"`
	TopHistory []*PlayHistory    `json:"top_histories"`
}

type World struct {
	ID   int64              `json:"world_id"`
	Name string             `json:"world_name"`
	Maps []*PlayerStatInMap `json:"maps"`
}

type RuleHistory struct {
	PlayHistoryID int64
	Rule          *MapRule
	IsPass        bool
	Value         int
}

type PlayHistory struct {
	ID              int64                        `json:"play_history_id"`
	GameSessionID   int64                        `json:"game_session_id"`
	ActionStep      int                          `json:"action_step"`
	NumberOfCommand int                          `json:"number_of_command"`
	IsFinited       bool                         `json:"is_finited"`
	IsCompleted     bool                         `json:"is_completed"`
	CommandMedal    sqlc_generated.NullMedalType `json:"command_medal"`
	ActionMedal     sqlc_generated.NullMedalType `json:"action_medal"`
	SubmitDatetime  time.Time                    `json:"submit_datetime"`
	Rules           []*RuleHistory               `json:"rules"`
}

type PlayerGameSession struct {
	GameSesssionID     int64          `json:"game_session_id"`
	PlayerID           string         `json:"player_id"`
	MapConfigurationID int64          `json:"map_configuration_id"`
	StartDatetime      time.Time      `json:"start_datetime"`
	EndDatetime        time.Time      `json:"end_datetime"`
	GameHistory        []*PlayHistory `json:"game_histories"`
}
