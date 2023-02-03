// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package sqlc_generated

import (
	"context"
)

type Querier interface {
	CreateLoginLog(ctx context.Context, playerID string) error
	CreateOrUpdatePlayer(ctx context.Context, arg CreateOrUpdatePlayerParams) error
	GetDoorFromMapConfigurationIDs(ctx context.Context, mapConfigurationIds []int64) ([]*GetDoorFromMapConfigurationIDsRow, error)
	GetItemFromMapConfigurationIDs(ctx context.Context, mapConfigurationIds []int64) ([]*GetItemFromMapConfigurationIDsRow, error)
	GetMapConfigFromPlayerID(ctx context.Context, playerID string) ([]*GetMapConfigFromPlayerIDRow, error)
	GetMapConfigurationWithItemFromPlayerID(ctx context.Context, playerID string) ([]*GetMapConfigurationWithItemFromPlayerIDRow, error)
	GetRuleFromMapConfigurationIDs(ctx context.Context, mapConfigurationIds []int64) ([]*MapConfigurationRule, error)
	ListWorldFromMapConfigurationIDs(ctx context.Context, mapConfigurationIds []int64) ([]*ListWorldFromMapConfigurationIDsRow, error)
}

var _ Querier = (*Queries)(nil)
