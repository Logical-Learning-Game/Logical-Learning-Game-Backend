// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: player.sql

package player

import (
	"context"
	"database/sql"
)

const createLoginLog = `-- name: CreateLoginLog :exec
INSERT INTO login_log (player_id)
VALUES (?)
`

func (q *Queries) CreateLoginLog(ctx context.Context, playerID string) error {
	_, err := q.db.ExecContext(ctx, createLoginLog, playerID)
	return err
}

const createOrUpdatePlayer = `-- name: CreateOrUpdatePlayer :exec
INSERT INTO player (player_id, email)
VALUES (?, ?)
ON DUPLICATE KEY UPDATE email = ?
`

type CreateOrUpdatePlayerParams struct {
	PlayerID string         `json:"player_id"`
	Email    sql.NullString `json:"email"`
}

func (q *Queries) CreateOrUpdatePlayer(ctx context.Context, arg CreateOrUpdatePlayerParams) error {
	_, err := q.db.ExecContext(ctx, createOrUpdatePlayer, arg.PlayerID, arg.Email, arg.Email)
	return err
}
