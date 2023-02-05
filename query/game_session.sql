-- name: CreateGameSession :one
INSERT INTO game_session (player_id, map_configuration_id, start_datetime, end_datetime)
VALUES ($1, $2, $3, $4)
RETURNING id, player_id, map_configuration_id, start_datetime, end_datetime;