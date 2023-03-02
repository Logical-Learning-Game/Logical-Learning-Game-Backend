package service

import (
	"context"
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
)

type PlayerStatisticService interface {
	CreateSessionHistory(ctx context.Context, playerID string, arg dto.CreateGameSessionRequestDTO) (*entity.GameSession, error)
	UpdateTopSubmitHistory(ctx context.Context, playerID string, args []*dto.TopSubmitHistoryDTO) ([]*entity.SubmitHistory, error)
	ListPlayerSessionData(ctx context.Context, playerID string) ([]*entity.GameSession, error)
}

type MapConfigurationService interface {
	ListPlayerAvailableMaps(ctx context.Context, playerID string) ([]*entity.World, error)
}
