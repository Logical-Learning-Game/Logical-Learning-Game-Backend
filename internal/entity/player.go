package entity

import "context"

type PlayerRepository interface {
	CreateLoginLog(ctx context.Context, playerID string) error
	CreateOrUpdatePlayer(ctx context.Context, playerID, email, name string) error
}

type PlayerService interface {
	CreateOrUpdatePlayerInformation(ctx context.Context, playerID, email, name string) error
	CreateLoginLog(ctx context.Context, playerID string) error
}
