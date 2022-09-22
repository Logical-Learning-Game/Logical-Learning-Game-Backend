-- name: CreateOrUpdatePlayer :exec
INSERT INTO player (player_id, email)
VALUES (?, sqlc.arg(email))
ON DUPLICATE KEY UPDATE email = sqlc.arg(email);

-- name: CreateLoginLog :exec
INSERT INTO login_log (player_id)
VALUES (?);