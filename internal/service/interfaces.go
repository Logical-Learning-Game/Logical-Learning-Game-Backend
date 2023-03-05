package service

import (
	"context"
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
)

type PlayerStatisticService interface {
	CreateSessionHistory(ctx context.Context, playerID string, arg dto.SessionHistoryDTO) (*entity.GameSession, error)
	UpdateTopSubmitHistory(ctx context.Context, playerID string, args []*dto.TopSubmitHistoryDTO) ([]*entity.SubmitHistory, error)
	ListPlayerSessionData(ctx context.Context, playerID string) ([]*dto.SessionHistoryDTO, error)
	ListTopSubmitHistory(ctx context.Context, playerID string) ([]*dto.TopSubmitHistoryDTO, error)
	GetPlayerData(ctx context.Context, playerID string) (*dto.SyncPlayerDataResponseDTO, error)
}

type MapConfigurationService interface {
	ListPlayerAvailableMaps(ctx context.Context, playerID string) ([]*dto.WorldDTO, error)
}

type PlayerService interface {
	LinkAccount(ctx context.Context, linkAccountRequestDTO dto.LinkAccountRequestDTO) (*entity.User, error)
	PlayerInfo(ctx context.Context, playerID string) (*dto.PlayerInfoResponseDTO, error)
	RemovePlayerData(ctx context.Context, playerID string) error
}
