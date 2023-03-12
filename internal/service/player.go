package service

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"

	"gorm.io/gorm"
)

type playerService struct {
	db *gorm.DB
}

func NewPlayerService(db *gorm.DB) PlayerService {
	return &playerService{
		db: db,
	}
}

var (
	defaultMaps = []int64{1, 2, 3, 4, 5, 6, 7}
)

func (s playerService) LinkAccount(ctx context.Context, linkAccountRequestDTO dto.LinkAccountRequest) (*entity.User, error) {
	user := &entity.User{
		PlayerID: linkAccountRequestDTO.PlayerID,
		Email:    linkAccountRequestDTO.Email,
	}

	txErr := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Create(user)
		if err := result.Error; err != nil {
			return err
		}

		initialDefaultMaps := make([]*entity.MapConfigurationForPlayer, 0, len(defaultMaps))
		for _, v := range defaultMaps {
			initialDefaultMaps = append(initialDefaultMaps, &entity.MapConfigurationForPlayer{
				PlayerID:           linkAccountRequestDTO.PlayerID,
				MapConfigurationID: v,
				IsPass:             false,
			})
		}

		result = tx.Create(&initialDefaultMaps)

		return result.Error
	})
	if txErr != nil {
		if pgErr, ok := txErr.(*pgconn.PgError); ok {
			switch pgErr.Code {
			case "23505":
				return nil, ErrAccountAlreadyLinked
			}
		} else {
			return nil, txErr
		}
	}

	return user, nil
}

func (s playerService) PlayerInfo(ctx context.Context, playerID string) (*dto.PlayerInfoResponse, error) {
	var user entity.User

	result := s.db.WithContext(ctx).
		Where(&entity.User{PlayerID: playerID}).
		First(&user)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPlayerNotFound
		} else {
			return nil, err
		}
	}

	playerInfoResponse := &dto.PlayerInfoResponse{
		PlayerID: user.PlayerID,
		Email:    user.Email,
	}

	return playerInfoResponse, nil
}

func (s playerService) RemovePlayerData(ctx context.Context, playerID string) error {
	txErr := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// remove all game session history for this player
		result := tx.Where(&entity.GameSession{
			PlayerID: playerID,
		}).Delete(&entity.GameSession{})
		if err := result.Error; err != nil {
			return err
		}

		// remove all top submit
		mapForPlayers := tx.Model(&entity.MapConfigurationForPlayer{}).
			Select("ID").
			Where(&entity.MapConfigurationForPlayer{
				PlayerID: playerID,
			})
		result = tx.Where("map_configuration_for_player_id IN (?)", mapForPlayers).
			Delete(&entity.SubmitHistory{})

		return result.Error
	})

	return txErr
}
