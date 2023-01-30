package repository

import (
	"context"
	"database/sql"
	"llg_backend/internal/entity"
	"strings"
)

type playerRepository struct {
	entity.Querier
}

func NewPlayerRepository(querier entity.Querier) PlayerRepository {
	return &playerRepository{
		Querier: querier,
	}
}

func (p playerRepository) CreateOrUpdatePlayer(ctx context.Context, playerID, email, name string) error {
	emailArg := sql.NullString{
		String: email,
		Valid:  true,
	}

	trimmedSpaceEmail := strings.TrimSpace(email)
	if trimmedSpaceEmail == "" {
		emailArg.Valid = false
	}

	nameArg := sql.NullString{
		String: name,
		Valid:  true,
	}

	trimmedSpaceName := strings.TrimSpace(name)
	if trimmedSpaceName == "" {
		nameArg.Valid = false
	}
	
	arg := entity.CreateOrUpdatePlayerParams{
		PlayerID: playerID,
		Email:    emailArg,
		Name:     nameArg,
	}

	return p.Querier.CreateOrUpdatePlayer(ctx, arg)
}