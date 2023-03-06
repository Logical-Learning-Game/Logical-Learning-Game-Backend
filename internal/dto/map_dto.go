package dto

import (
	"llg_backend/internal/entity"
	"llg_backend/internal/entity/nullable"
)

type WorldDTO struct {
	WorldID   int64                  `json:"world_id"`
	WorldName string                 `json:"world_name"`
	Maps      []*MapConfigurationDTO `json:"maps"`
}

type MapConfigurationDTO struct {
	MapID                      int64               `json:"map_id"`
	MapName                    string              `json:"map_name"`
	Tile                       []int               `json:"tile"`
	Height                     int                 `json:"height"`
	Width                      int                 `json:"width"`
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
}

type RuleDTO struct {
	MapRuleID  int64            `json:"map_rule_id"`
	RuleName   string           `json:"rule_name"`
	Theme      entity.RuleTheme `json:"rule_theme"`
	Parameters []int            `json:"parameters"`
}
