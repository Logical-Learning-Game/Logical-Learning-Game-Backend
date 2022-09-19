-- name: CreatePlayer :exec
INSERT INTO player (player_id, email)
VALUES (?, ?);
