package service

import (
	"context"
)

type PlayerService interface {
	CreateOrUpdatePlayerInformation(ctx context.Context, playerID, email, name string) error
	CreateLoginLog(ctx context.Context, playerID string) error
}
