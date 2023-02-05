// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: play_history.sql

package sqlc_generated

import (
	"context"
	"time"
)

const createPlayHistory = `-- name: CreatePlayHistory :one
INSERT INTO play_history (game_session_id, action_step, number_of_command, is_finited, is_completed, submit_datetime)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, game_session_id, action_step, number_of_command, is_finited, is_completed, submit_datetime
`

type CreatePlayHistoryParams struct {
	GameSessionID   int64     `json:"game_session_id"`
	ActionStep      int32     `json:"action_step"`
	NumberOfCommand int32     `json:"number_of_command"`
	IsFinited       bool      `json:"is_finited"`
	IsCompleted     bool      `json:"is_completed"`
	SubmitDatetime  time.Time `json:"submit_datetime"`
}

func (q *Queries) CreatePlayHistory(ctx context.Context, arg CreatePlayHistoryParams) (*PlayHistory, error) {
	row := q.db.QueryRowContext(ctx, createPlayHistory,
		arg.GameSessionID,
		arg.ActionStep,
		arg.NumberOfCommand,
		arg.IsFinited,
		arg.IsCompleted,
		arg.SubmitDatetime,
	)
	var i PlayHistory
	err := row.Scan(
		&i.ID,
		&i.GameSessionID,
		&i.ActionStep,
		&i.NumberOfCommand,
		&i.IsFinited,
		&i.IsCompleted,
		&i.SubmitDatetime,
	)
	return &i, err
}
