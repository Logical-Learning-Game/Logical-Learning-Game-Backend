-- name: CreatePlayHistory :one
INSERT INTO play_history (game_session_id, action_step, number_of_command, is_finited, is_completed, command_medal,
                          action_medal, submit_datetime)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, game_session_id, action_step, number_of_command, is_finited, is_completed, command_medal, action_medal, submit_datetime;

-- name: CreateRuleHistory :one
INSERT INTO play_history_rule (play_history_id, map_configuration_rule_id, is_pass)
VALUES ($1, $2, $3)
RETURNING play_history_id, map_configuration_rule_id, is_pass;

-- name: CreateStateValue :one
INSERT INTO state_value (play_history_id, command_count, forward_command_count, right_command_count, back_command_count,
                         left_command_count, condition_command_count, action_count, forward_action_count,
                         right_action_count, back_action_count, left_action_count, condition_action_count)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
RETURNING play_history_id, command_count, forward_command_count, right_command_count, back_command_count, left_command_count, condition_command_count, action_count, forward_action_count, right_action_count, back_action_count, left_action_count, condition_action_count;

-- name: CreateCommandNode :one
INSERT INTO command_node (play_history_id, type, in_game_position_x, in_game_position_y)
VALUES ($1, $2, $3, $4)
RETURNING id, play_history_id, type, in_game_position_x, in_game_position_y;

-- name: CreateCommandEdge :one
INSERT INTO command_edge (source_node_id, destination_node_id, type)
VALUES ($1, $2, $3)
RETURNING source_node_id, destination_node_id, type;