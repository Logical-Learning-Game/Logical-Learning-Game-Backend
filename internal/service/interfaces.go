package service

import (
	"context"
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
)

type PlayerStatisticService interface {
	CreateSessionHistory(ctx context.Context, playerID string, arg dto.CreateGameSessionRequestDTO) (*entity.GameSession, error)
	UpdateTopSubmitHistory(ctx context.Context, playerID string, args []*dto.TopSubmitHistoryDTO) ([]*entity.SubmitHistory, error)
}

type MapConfigurationService interface {
	ListFromPlayerID(ctx context.Context, playerID string) ([]*entity.World, error)
}
