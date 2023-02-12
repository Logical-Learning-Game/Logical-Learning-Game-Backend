package repository

import (
	"context"
	"database/sql"
	"llg_backend/internal/entity"
	"llg_backend/internal/entity/sqlc_generated"
	"strings"
)

type playerRepository struct {
	sqlc_generated.Querier
}

func NewPlayerRepository(querier sqlc_generated.Querier) entity.PlayerRepository {
	return &playerRepository{
		Querier: querier,
	}
}

func (p *playerRepository) CreateOrUpdatePlayer(ctx context.Context, playerID, email, name string) error {
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

	arg := sqlc_generated.CreateOrUpdatePlayerParams{
		PlayerID: playerID,
		Email:    emailArg,
		Name:     nameArg,
	}

	return p.Querier.CreateOrUpdatePlayer(ctx, arg)
}
