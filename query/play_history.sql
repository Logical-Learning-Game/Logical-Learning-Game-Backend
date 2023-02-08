-- name: CreatePlayHistory :one
INSERT INTO play_history (game_session_id, action_step, number_of_command, is_finited, is_completed, command_medal,
                          action_medal, submit_datetime)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, game_session_id, action_step, number_of_command, is_finited, is_completed, command_medal, action_medal, submit_datetime;