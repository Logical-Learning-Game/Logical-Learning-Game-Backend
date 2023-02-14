// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: map_configuration.sql

package sqlc_generated

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const getMapConfigFromPlayerID = `-- name: GetMapConfigFromPlayerID :many
SELECT map_conf.id     AS map_config_id,
       map_conf.world_id,
       world.name      AS world_name,
       map_conf.tile_array,
       map_conf.height AS map_height,
       map_conf.width  AS map_width,
       map_conf.start_player_direction,
       map_conf.start_player_position_x,
       map_conf.start_player_position_y,
       map_conf.goal_position_x,
       map_conf.goal_position_y,
       map_conf.config_name,
       map_conf.map_image_path,
       map_conf.difficulty,
       map_conf.star_requirement,
       map_conf.least_solvable_command_gold,
       map_conf.least_solvable_command_silver,
       map_conf.least_solvable_command_bronze,
       map_conf.least_solvable_action_gold,
       map_conf.least_solvable_action_silver,
       map_conf.least_solvable_action_bronze,
       map_player.is_pass
FROM map_configuration_for_player AS map_player
         INNER JOIN map_configuration AS map_conf ON map_player.map_configuration_id = map_conf.id
         INNER JOIN world ON world.id = map_conf.world_id
WHERE player_id = $1
`

type GetMapConfigFromPlayerIDRow struct {
	MapConfigID                int64          `json:"map_config_id"`
	WorldID                    int64          `json:"world_id"`
	WorldName                  string         `json:"world_name"`
	TileArray                  []int32        `json:"tile_array"`
	MapHeight                  int32          `json:"map_height"`
	MapWidth                   int32          `json:"map_width"`
	StartPlayerDirection       MapDirection   `json:"start_player_direction"`
	StartPlayerPositionX       int32          `json:"start_player_position_x"`
	StartPlayerPositionY       int32          `json:"start_player_position_y"`
	GoalPositionX              int32          `json:"goal_position_x"`
	GoalPositionY              int32          `json:"goal_position_y"`
	ConfigName                 string         `json:"config_name"`
	MapImagePath               sql.NullString `json:"map_image_path"`
	Difficulty                 MapDifficulty  `json:"difficulty"`
	StarRequirement            int32          `json:"star_requirement"`
	LeastSolvableCommandGold   int32          `json:"least_solvable_command_gold"`
	LeastSolvableCommandSilver int32          `json:"least_solvable_command_silver"`
	LeastSolvableCommandBronze int32          `json:"least_solvable_command_bronze"`
	LeastSolvableActionGold    int32          `json:"least_solvable_action_gold"`
	LeastSolvableActionSilver  int32          `json:"least_solvable_action_silver"`
	LeastSolvableActionBronze  int32          `json:"least_solvable_action_bronze"`
	IsPass                     bool           `json:"is_pass"`
}

func (q *Queries) GetMapConfigFromPlayerID(ctx context.Context, playerID string) ([]*GetMapConfigFromPlayerIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getMapConfigFromPlayerID, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetMapConfigFromPlayerIDRow{}
	for rows.Next() {
		var i GetMapConfigFromPlayerIDRow
		if err := rows.Scan(
			&i.MapConfigID,
			&i.WorldID,
			&i.WorldName,
			pq.Array(&i.TileArray),
			&i.MapHeight,
			&i.MapWidth,
			&i.StartPlayerDirection,
			&i.StartPlayerPositionX,
			&i.StartPlayerPositionY,
			&i.GoalPositionX,
			&i.GoalPositionY,
			&i.ConfigName,
			&i.MapImagePath,
			&i.Difficulty,
			&i.StarRequirement,
			&i.LeastSolvableCommandGold,
			&i.LeastSolvableCommandSilver,
			&i.LeastSolvableCommandBronze,
			&i.LeastSolvableActionGold,
			&i.LeastSolvableActionSilver,
			&i.LeastSolvableActionBronze,
			&i.IsPass,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
