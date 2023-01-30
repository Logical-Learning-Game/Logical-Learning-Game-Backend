package service

import (
	"context"
	"llg_backend/internal/repository"
	"llg_backend/pkg/logger"
	"strings"
)

type playerService struct {
	playerRepo repository.PlayerRepository
}

func NewPlayerService(playerRepo repository.PlayerRepository) PlayerService {
	return &playerService{
		playerRepo: playerRepo,
	}
}

func (s *playerService) CreateOrUpdatePlayerInformation(ctx context.Context, playerID, email, name string) error {
	return s.playerRepo.CreateOrUpdatePlayer(ctx, playerID, email, name)
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
	sanitizedPlayerID := strings.Replace(playerID, "\n", "", -1)
	sanitizedPlayerID = strings.Replace(sanitizedPlayerID, "\r", "", -1)
	s.log.Debugw("CreateLoginLog - param", "playerID", sanitizedPlayerID)

	err := s.PlayerService.CreateLoginLog(ctx, playerID)
	if err != nil {
		s.log.Errorw("CreateLoginLog error", "err", err)
	}

	s.log.Debugw("CreateLoginLog - return", "err", err)
	return err
}
