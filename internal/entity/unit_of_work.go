package entity

import "context"

type UnitOfWorkStore struct {
	DoorRepo        DoorRepository
	GameSessionRepo GameSessionRepository
	ItemRepo        ItemRepository
	MapConfigRepo   MapConfigurationRepository
	PlayHistoryRepo PlayHistoryRepository
	PlayerRepo      PlayerRepository
	RuleRepo        RuleRepository
	WorldRepo       WorldRepository
}

type UnitOfWork interface {
	Do(ctx context.Context, fn UnitOfWorkBlock) error
}

type UnitOfWorkBlock func(*UnitOfWorkStore) error
