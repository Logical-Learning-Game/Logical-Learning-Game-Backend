package service

import (
	"context"
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
	"time"
)

type PlayerStatisticService interface {
	CreateSessionHistory(ctx context.Context, playerID string, arg dto.SessionHistoryRequest) (*entity.GameSession, error)
	UpdateTopSubmitHistory(ctx context.Context, playerID string, args []*dto.TopSubmitHistoryRequest) ([]*entity.SubmitHistory, error)
	ListPlayerSessionDataForGame(ctx context.Context, playerID string) ([]*dto.SessionHistoryForGameResponse, error)
	ListTopSubmitHistory(ctx context.Context, playerID string) ([]*dto.TopSubmitHistoryResponse, error)
	GetPlayerData(ctx context.Context, playerID string) (*dto.PlayerDataDTO, error)

	ListPlayerSessionForAdmin(ctx context.Context, playerID string) ([]*dto.SessionDataForAdminResponse, error)
	ListSubmitHistoriesForAdmin(ctx context.Context, sessionID int64) ([]*dto.SubmitHistoryForAdminResponse, error)
	ListMapOfPlayerInfoForAdmin(ctx context.Context, playerID string) ([]*dto.MapOfPlayerInfoForAdminResponse, error)
	ListPlayerSignInHistory(ctx context.Context, playerID string) ([]time.Time, error)
}

type MapConfigurationService interface {
	ListPlayerAvailableMaps(ctx context.Context, playerID string) ([]*dto.WorldDTO, error)

	ListWorld(ctx context.Context) ([]*dto.WorldForAdminResponse, error)
	CreateWorld(ctx context.Context, name string) error
	UpdateWorld(ctx context.Context, worldID int64, name string) error
	ListWorldWithMap(ctx context.Context) ([]*dto.WorldWithMapForAdminResponse, error)
	UpdateMapOfPlayerActive(ctx context.Context, playerID string, mapID int64, active bool) error
	SetMapActive(ctx context.Context, mapID int64, active bool) error
	GetMapByID(ctx context.Context, mapID int64) (*dto.MapConfigurationForAdminDTO, error)
	CreateMap(ctx context.Context, createMapRequest *dto.CreateMapRequest, imagePath string) error
	UpdateMap(ctx context.Context, mapID int64, createMapRequest *dto.CreateMapRequest, imagePath string) error
	AddMapToAllPlayers(ctx context.Context, mapID int64) error
}

type PlayerService interface {
	LinkAccount(ctx context.Context, linkAccountRequestDTO dto.LinkAccountRequest) (*entity.User, error)
	PlayerInfo(ctx context.Context, playerID string) (*dto.PlayerInfoResponse, error)
	RemovePlayerData(ctx context.Context, playerID string) error
	ListPlayers(ctx context.Context) ([]*dto.PlayerInfoResponse, error)
}

type AdminAuthenticationService interface {
	Login(ctx context.Context, username, password string) (*dto.AdminLoginResponse, error)
}
