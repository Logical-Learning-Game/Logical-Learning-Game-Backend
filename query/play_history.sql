-- name: CreatePlayHistory :one
INSERT INTO play_history (game_session_id, action_step, number_of_command, is_finited, is_completed, submit_datetime)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, game_session_id, action_step, number_of_command, is_finited, is_completed, submit_datetime;