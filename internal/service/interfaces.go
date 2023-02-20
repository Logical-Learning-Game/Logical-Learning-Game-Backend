package service

import (
	"context"
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
)

type PlayerStatisticService interface {
	CreateSessionHistory(ctx context.Context, playerID string, arg dto.CreateGameSessionRequestDTO) (*entity.GameSession, error)
}

type MapConfigurationService interface {
	ListFromPlayerID(ctx context.Context, playerID string) ([]*entity.World, error)
}
