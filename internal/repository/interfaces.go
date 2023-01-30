package repository

import (
	"context"
)

type PlayerRepository interface {
	CreateLoginLog(ctx context.Context, playerID string) error
	CreateOrUpdatePlayer(ctx context.Context, playerID, email, name string) error
}
