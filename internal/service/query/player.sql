-- name: CreateOrUpdatePlayer :exec
INSERT INTO player (player_id, email, name)
VALUES (?, sqlc.arg(email), sqlc.arg(name))
ON DUPLICATE KEY UPDATE email = sqlc.arg(email),
                        name  = sqlc.arg(name);

-- name: CreateLoginLog :exec
INSERT INTO login_log (player_id)
VALUES (?);