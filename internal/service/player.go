package service

import (
	"context"
	"llg_backend/internal/service/repository/player"
	"llg_backend/pkg/logger"
)

type PlayerService interface {
	CreateOrUpdatePlayer(ctx context.Context, arg player.CreateOrUpdatePlayerParams) error
	CreateLoginLog(ctx context.Context, playerID string) error
}

type playerService struct {
	playerRepo player.Querier
}

func NewPlayerService(playerRepo player.Querier) PlayerService {
	return &playerService{
		playerRepo: playerRepo,
	}
}

func (s *playerService) CreateOrUpdatePlayer(ctx context.Context, arg player.CreateOrUpdatePlayerParams) error {
	return s.playerRepo.CreateOrUpdatePlayer(ctx, arg)
}

func (s *playerService) CreateLoginLog(ctx context.Context, playerID string) error {
	return s.playerRepo.CreateLoginLog(ctx, playerID)
}

type playerServiceWithLog struct {
	PlayerService

	log logger.Logger
}

func NewPlayerServiceWithLog(playerService PlayerService, log logger.Logger) PlayerService {
	return &playerServiceWithLog{
		PlayerService: playerService,
		log:           log,
	}
}

func (s *playerServiceWithLog) CreateLoginLog(ctx context.Context, playerID string) error {
	s.log.Debugw("CreateLoginLog - param", "playerID", playerID)
	err := s.PlayerService.CreateLoginLog(ctx, playerID)
	if err != nil {
		s.log.Errorw("CreateLoginLog error", "err", err)
	}
	s.log.Debugw("CreateLoginLog - return", "err", err)
	return err
}
