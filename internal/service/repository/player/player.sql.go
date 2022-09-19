// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: player.sql

package player

import (
	"context"
)

const createPlayer = `-- name: CreatePlayer :exec
INSERT INTO player (player_id, email)
VALUES (?, ?)
`

type CreatePlayerParams struct {
	PlayerID string `json:"player_id"`
	Email    string `json:"email"`
}

func (q *Queries) CreatePlayer(ctx context.Context, arg CreatePlayerParams) error {
	_, err := q.db.ExecContext(ctx, createPlayer, arg.PlayerID, arg.Email)
	return err
}
