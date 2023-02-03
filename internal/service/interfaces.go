package service

import (
	"context"
	"llg_backend/internal/entity"
)

type PlayerService interface {
	CreateOrUpdatePlayerInformation(ctx context.Context, playerID, email, name string) error
	CreateLoginLog(ctx context.Context, playerID string) error
}

type MapConfigurationService interface {
	ListFromPlayerID(ctx context.Context, playerID string) ([]*entity.PlayerStatInMap, error)
}

type WorldService interface {
	ListFromPlayerID(ctx context.Context, playerID string) ([]*entity.World, error)
}
