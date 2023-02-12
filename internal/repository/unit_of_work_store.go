package repository

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
