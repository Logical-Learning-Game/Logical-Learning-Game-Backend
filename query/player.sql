-- name: CreateOrUpdatePlayer :exec
INSERT INTO player (player_id, email, name)
VALUES ($1, $2, $3)
ON CONFLICT (player_id)
    DO UPDATE SET
    email = $2,
    name = $3;