-- name: CreateLoginLog :exec
INSERT INTO login_log (player_id)
VALUES ($1);