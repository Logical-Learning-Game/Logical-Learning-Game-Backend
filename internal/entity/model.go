package entity

import (
	"database/sql"
	"github.com/lib/pq"
	"llg_backend/internal/entity/nullable"
	"time"
)

type Vector2Int struct {
	X int32 `gorm:"not null"`
	Y int32 `gorm:"not null"`
}

type Vector2Float32 struct {
	X float32 `gorm:"not null"`
	Y float32 `gorm:"not null"`
}

type Admin struct {
	Username       string `gorm:"primaryKey;type:varchar(255)"`
	HashedPassword string `gorm:"type:varchar(255);not null"`
}

type User struct {
	PlayerID          string                       `gorm:"primaryKey;type:varchar(255)"`
	Email             string                       `gorm:"type:varchar(255)"`
	SignInHistories   []*SignInHistory             `gorm:"foreignKey:PlayerID"`
	MapConfigurations []*MapConfigurationForPlayer `gorm:"foreignKey:PlayerID"`
	GameSession       []*GameSession               `gorm:"foreignKey:PlayerID"`
}

type SignInHistory struct {
	PlayerID  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
}

type Item struct {
	ID     int64    `gorm:"primaryKey"`
	Name   string   `gorm:"type:varchar(255);not null"`
	Type   ItemType `gorm:"type:item_type;not null"`
	Active bool     `gorm:"not null"`
}

type Door struct {
	ID     int64    `gorm:"primaryKey"`
	Name   string   `gorm:"type:varchar(255);not null"`
	Type   DoorType `gorm:"type:door_type;not null"`
	Active bool     `gorm:"not null"`
}

type Rule struct {
	Name   string `gorm:"primaryKey;type:varchar(255)"`
	Active bool   `gorm:"not null"`
}

type World struct {
	ID                int64               `gorm:"primaryKey"`
	Name              string              `gorm:"type:varchar(255);not null"`
	MapConfigurations []*MapConfiguration `gorm:"foreignKey:WorldID"`
}

type MapConfiguration struct {
	ID                         int64                   `gorm:"primaryKey"`
	WorldID                    int64                   `gorm:"not null"`
	World                      *World                  `gorm:"foreignKey:WorldID"`
	ConfigName                 string                  `gorm:"type:varchar(255);not null"`
	Tile                       pq.Int32Array           `gorm:"type:integer[];not null"`
	Height                     int32                   `gorm:"not null"`
	Width                      int32                   `gorm:"not null"`
	StartPlayerDirection       MapDirection            `gorm:"type:map_direction;not null"`
	StartPlayerPosition        Vector2Int              `gorm:"embedded;embeddedPrefix:start_player_position_"`
	GoalPosition               Vector2Int              `gorm:"embedded;embeddedPrefix:goal_position_"`
	MapImagePath               sql.NullString          `gorm:"type:varchar(255)"`
	Difficulty                 Difficulty              `gorm:"type:difficulty;not null"`
	StarRequirement            int32                   `gorm:"not null"`
	LeastSolvableCommandGold   int32                   `gorm:"not null"`
	LeastSolvableCommandSilver int32                   `gorm:"not null"`
	LeastSolvableCommandBronze int32                   `gorm:"not null"`
	LeastSolvableActionGold    int32                   `gorm:"not null"`
	LeastSolvableActionSilver  int32                   `gorm:"not null"`
	LeastSolvableActionBronze  int32                   `gorm:"not null"`
	CreatedAt                  time.Time               `gorm:"not null;default:now()"`
	Active                     bool                    `gorm:"not null;default:true"`
	Items                      []*MapConfigurationItem `gorm:"foreignKey:MapConfigurationID"`
	Rules                      []*MapConfigurationRule `gorm:"foreignKey:MapConfigurationID"`
	Doors                      []*MapConfigurationDoor `gorm:"foreignKey:MapConfigurationID"`
}

type MapConfigurationItem struct {
	ID                 int64      `gorm:"primaryKey"`
	MapConfigurationID int64      `gorm:"not null"`
	ItemID             int64      `gorm:"not null"`
	Item               *Item      `gorm:"foreignKey:ItemID"`
	Position           Vector2Int `gorm:"embedded;embeddedPrefix:position_"`
}

type MapConfigurationRule struct {
	ID                 int64         `gorm:"primaryKey"`
	MapConfigurationID int64         `gorm:"not null"`
	RuleName           string        `gorm:"not null"`
	Rule               *Rule         `gorm:"foreignKey:RuleName"`
	Theme              RuleTheme     `gorm:"type:rule_theme;not null"`
	Parameters         pq.Int32Array `gorm:"type:integer[];not null"`
}

type MapConfigurationDoor struct {
	ID                 int64        `gorm:"primaryKey"`
	MapConfigurationID int64        `gorm:"not null"`
	DoorID             int64        `gorm:"not null"`
	Door               *Door        `gorm:"foreignKey:DoorID"`
	Position           Vector2Int   `gorm:"embedded;embeddedPrefix:position_"`
	DoorDirection      MapDirection `gorm:"type:map_direction;not null"`
}

type MapConfigurationForPlayer struct {
	ID                 int64             `gorm:"primaryKey"`
	PlayerID           string            `gorm:"not null"`
	MapConfigurationID int64             `gorm:"not null"`
	MapConfiguration   *MapConfiguration `gorm:"foreignKey:MapConfigurationID"`
	IsPass             bool              `gorm:"not null;default:false"`
	Active             bool              `gorm:"not null;default:true"`
	TopSubmitHistory   *SubmitHistory    `gorm:"foreignKey:MapConfigurationForPlayerID;constraint:OnDelete:CASCADE"`
}

type GameSession struct {
	ID                 int64             `gorm:"primaryKey"`
	PlayerID           string            `gorm:"type:varchar(255);not null"`
	MapConfigurationID int64             `gorm:"not null"`
	MapConfiguration   *MapConfiguration `gorm:"foreignKey:MapConfigurationID"`
	StartDatetime      time.Time         `gorm:"not null"`
	EndDatetime        nullable.NullTime
	SubmitHistories    []*SubmitHistory `gorm:"foreignKey:GameSessionID;constraint:OnDelete:CASCADE"`
}

type SubmitHistory struct {
	ID                          int64 `gorm:"primaryKey"`
	GameSessionID               int64
	MapConfigurationForPlayerID int64
	IsFinited                   bool                 `gorm:"not null"`
	IsCompleted                 bool                 `gorm:"not null"`
	CommandMedal                MedalType            `gorm:"type:medal_type;not null"`
	ActionMedal                 MedalType            `gorm:"type:medal_type;not null"`
	SubmitDatetime              time.Time            `gorm:"not null"`
	StateValue                  *StateValue          `gorm:"foreignKey:SubmitHistoryID;constraint:OnDelete:CASCADE"`
	SubmitHistoryRules          []*SubmitHistoryRule `gorm:"foreignKey:SubmitHistoryID;constraint:OnDelete:CASCADE"`
	CommandNodes                []*CommandNode       `gorm:"foreignKey:SubmitHistoryID;constraint:OnDelete:CASCADE"`
	CommandEdges                []*CommandEdge       `gorm:"foreignKey:SubmitHistoryID;constraint:OnDelete:CASCADE"`
}

type StateValue struct {
	SubmitHistoryID       int64 `gorm:"not null"`
	CommandCount          int32 `gorm:"not null"`
	ForwardCommandCount   int32 `gorm:"not null"`
	RightCommandCount     int32 `gorm:"not null"`
	BackCommandCount      int32 `gorm:"not null"`
	LeftCommandCount      int32 `gorm:"not null"`
	ConditionCommandCount int32 `gorm:"not null"`
	ActionCount           int32 `gorm:"not null"`
	ForwardActionCount    int32 `gorm:"not null"`
	RightActionCount      int32 `gorm:"not null"`
	BackActionCount       int32 `gorm:"not null"`
	LeftActionCount       int32 `gorm:"not null"`
	ConditionActionCount  int32 `gorm:"not null"`
	AllItemCount          int32 `gorm:"not null"`
	KeyACount             int32 `gorm:"not null"`
	KeyBCount             int32 `gorm:"not null"`
	KeyCCount             int32 `gorm:"not null"`
}

type SubmitHistoryRule struct {
	SubmitHistoryID        int64                 `gorm:"not null"`
	MapConfigurationRuleID int64                 `gorm:"not null"`
	MapConfigurationRule   *MapConfigurationRule `gorm:"foreignKey:MapConfigurationRuleID"`
	IsPass                 bool                  `gorm:"not null"`
}

type CommandNode struct {
	SubmitHistoryID int64           `gorm:"not null"`
	Index           int32           `gorm:"not null"`
	Type            CommandNodeType `gorm:"type:command_node_type;not null"`
	InGamePosition  Vector2Float32  `gorm:"embedded;embeddedPrefix:in_game_position_"`
}

type CommandEdge struct {
	SubmitHistoryID      int64           `gorm:"not null"`
	SourceNodeIndex      int32           `gorm:"not null"`
	DestinationNodeIndex int32           `gorm:"not null"`
	Type                 CommandEdgeType `gorm:"type:command_edge_type;not null"`
}
