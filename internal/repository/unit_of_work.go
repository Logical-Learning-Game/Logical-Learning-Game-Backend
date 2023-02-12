package repository

import (
	"context"
	"database/sql"
	"fmt"
	"llg_backend/internal/entity/sqlc_generated"
)

type unitOfWork struct {
	db *sql.DB
}

func NewUnitOfWork(db *sql.DB) UnitOfWork {
	return &unitOfWork{db: db}
}

func (u unitOfWork) Do(ctx context.Context, fn UnitOfWorkBlock) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := sqlc_generated.New(tx)

	uows := &UnitOfWorkStore{
		DoorRepo:        NewDoorRepository(q),
		GameSessionRepo: NewGameSessionRepository(q),
		ItemRepo:        NewItemRepository(q),
		MapConfigRepo:   NewMapConfigurationRepository(q),
		PlayHistoryRepo: NewPlayHistoryRepository(q),
		PlayerRepo:      NewPlayerRepository(q),
		RuleRepo:        NewRuleRepository(q),
		WorldRepo:       NewWorldRepository(q),
	}

	if err = fn(uows); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
