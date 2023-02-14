-- name: GetMapConfigFromPlayerID :many
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
WHERE player_id = $1;