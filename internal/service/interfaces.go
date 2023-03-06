package service

import (
	"context"
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
)

type PlayerStatisticService interface {
	CreateSessionHistory(ctx context.Context, playerID string, arg dto.SessionHistoryRequest) (*entity.GameSession, error)
	UpdateTopSubmitHistory(ctx context.Context, playerID string, args []*dto.TopSubmitHistoryRequest) ([]*entity.SubmitHistory, error)
	ListPlayerSessionData(ctx context.Context, playerID string) ([]*dto.SessionHistoryResponse, error)
	ListTopSubmitHistory(ctx context.Context, playerID string) ([]*dto.TopSubmitHistoryResponse, error)
	GetPlayerData(ctx context.Context, playerID string) (*dto.PlayerDataDTO, error)
}

type MapConfigurationService interface {
	ListPlayerAvailableMaps(ctx context.Context, playerID string) ([]*dto.WorldDTO, error)
}

type PlayerService interface {
	LinkAccount(ctx context.Context, linkAccountRequestDTO dto.LinkAccountRequest) (*entity.User, error)
	PlayerInfo(ctx context.Context, playerID string) (*dto.PlayerInfoResponse, error)
	RemovePlayerData(ctx context.Context, playerID string) error
}
