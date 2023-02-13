package entity

import (
	"context"
	"llg_backend/internal/entity/sqlc_generated"
)

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
	WorldID              int64                        `json:"-"`
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

type MapConfigurationRepository interface {
	ListFromPlayerID(ctx context.Context, playerID string) ([]*MapConfiguration, error)
}

type MapConfigurationService interface {
	ListFromPlayerID(ctx context.Context, playerID string) ([]*MapConfiguration, error)
}

type MapItem struct {
	MapConfigID int64                   `json:"-"`
	Name        string                  `json:"name"`
	Type        sqlc_generated.ItemType `json:"type"`
	Position    Vector2Int              `json:"position"`
}

type ItemRepository interface {
	ListFromMapConfigurationIDs(ctx context.Context, mapConfigurationIDs []int64) ([]*MapItem, error)
}

type MapDoor struct {
	MapConfigID   int64                       `json:"-"`
	Name          string                      `json:"name"`
	Type          sqlc_generated.DoorType     `json:"type"`
	Position      Vector2Int                  `json:"position"`
	DoorDirection sqlc_generated.MapDirection `json:"door_direction"`
}

type DoorRepository interface {
	ListFromMapConfigurationIDs(ctx context.Context, mapConfigurationIDs []int64) ([]*MapDoor, error)
}

type MapRule struct {
	ID          int64                    `json:"map_config_rule_id"`
	MapConfigID int64                    `json:"-"`
	Type        string                   `json:"type"`
	Order       int                      `json:"order"`
	Theme       sqlc_generated.RuleTheme `json:"theme"`
	Parameters  []int                    `json:"parameters"`
}

type RuleRepository interface {
	ListFromMapConfigurationIDs(ctx context.Context, mapConfigurationIDs []int64) ([]*MapRule, error)
}

type World struct {
	ID   int64               `json:"world_id"`
	Name string              `json:"world_name"`
	Maps []*MapConfiguration `json:"maps"`
}

type WorldRepository interface {
	ListFromMapConfigurationIDs(ctx context.Context, mapConfigurationIDs []int64) ([]*World, error)
}

type WorldService interface {
	ListFromPlayerID(ctx context.Context, playerID string) ([]*World, error)
}
