package entity

import (
	"database/sql/driver"
	"fmt"
)

type ItemType string

const (
	ItemTypeKeyA ItemType = "key_a"
	ItemTypeKeyB ItemType = "key_b"
	ItemTypeKeyC ItemType = "key_c"
)

func (i *ItemType) Scan(value interface{}) error {
	switch t := value.(type) {
	case []byte:
		*i = ItemType(t)
	case string:
		*i = ItemType(t)
	default:
		return fmt.Errorf("unsupported scan type for ItemType: %T", t)
	}
	return nil
}

func (i ItemType) Value() (driver.Value, error) {
	switch i {
	case ItemTypeKeyA, ItemTypeKeyB, ItemTypeKeyC:
		return string(i), nil
	default:
		return nil, fmt.Errorf("invalid ItemType value: %v", i)
	}
}

type DoorType string

const (
	DoorTypeNoKey DoorType = "door_no_key"
	DoorTypeA     DoorType = "door_a"
	DoorTypeB     DoorType = "door_b"
	DoorTypeC     DoorType = "door_c"
)

func (i *DoorType) Scan(value interface{}) error {
	switch t := value.(type) {
	case []byte:
		*i = DoorType(t)
	case string:
		*i = DoorType(t)
	default:
		return fmt.Errorf("unsupported scan type for DoorType: %T", t)
	}
	return nil
}

func (i DoorType) Value() (driver.Value, error) {
	switch i {
	case DoorTypeNoKey, DoorTypeA, DoorTypeB, DoorTypeC:
		return string(i), nil
	default:
		return nil, fmt.Errorf("invalid DoorType value: %v", i)
	}
}

type MapDirection string

const (
	MapDirectionUp    MapDirection = "up"
	MapDirectionRight MapDirection = "right"
	MapDirectionDown  MapDirection = "down"
	MapDirectionLeft  MapDirection = "left"
)

func (i *MapDirection) Scan(value interface{}) error {
	switch t := value.(type) {
	case []byte:
		*i = MapDirection(t)
	case string:
		*i = MapDirection(t)
	default:
		return fmt.Errorf("unsupported scan type for MapDirection: %T", t)
	}
	return nil
}

func (i MapDirection) Value() (driver.Value, error) {
	switch i {
	case MapDirectionUp, MapDirectionRight, MapDirectionDown, MapDirectionLeft:
		return string(i), nil
	default:
		return nil, fmt.Errorf("invalid MapDirection value: %v", i)
	}
}

type Difficulty string

const (
	DifficultyEasy   Difficulty = "easy"
	DifficultyMedium Difficulty = "medium"
	DifficultyHard   Difficulty = "hard"
)

func (i *Difficulty) Scan(value interface{}) error {
	switch t := value.(type) {
	case []byte:
		*i = Difficulty(t)
	case string:
		*i = Difficulty(t)
	default:
		return fmt.Errorf("unsupported scan type for Difficulty: %T", t)
	}
	return nil
}

func (i Difficulty) Value() (driver.Value, error) {
	switch i {
	case DifficultyEasy, DifficultyMedium, DifficultyHard:
		return string(i), nil
	default:
		return nil, fmt.Errorf("invalid Difficulty value: %v", i)
	}
}

type RuleTheme string

const (
	RuleThemeNormal      RuleTheme = "normal"
	RuleThemeConditional RuleTheme = "conditional"
	RuleThemeLoop        RuleTheme = "loop"
)

func (i *RuleTheme) Scan(value interface{}) error {
	switch t := value.(type) {
	case []byte:
		*i = RuleTheme(t)
	case string:
		*i = RuleTheme(t)
	default:
		return fmt.Errorf("unsupported scan type for RuleTheme: %T", t)
	}
	return nil
}

func (i RuleTheme) Value() (driver.Value, error) {
	switch i {
	case RuleThemeNormal, RuleThemeConditional, RuleThemeLoop:
		return string(i), nil
	default:
		return nil, fmt.Errorf("invalid RuleTheme value: %v", i)
	}
}

type MedalType string

const (
	MedalTypeGold   MedalType = "gold"
	MedalTypeSilver MedalType = "silver"
	MedalTypeBronze MedalType = "bronze"
	MedalTypeNone   MedalType = "none"
)

func (i *MedalType) Scan(value interface{}) error {
	switch t := value.(type) {
	case []byte:
		*i = MedalType(t)
	case string:
		*i = MedalType(t)
	default:
		return fmt.Errorf("unsupported scan type for MedalType: %T", t)
	}
	return nil
}

func (i MedalType) Value() (driver.Value, error) {
	switch i {
	case MedalTypeGold, MedalTypeSilver, MedalTypeBronze, MedalTypeNone:
		return string(i), nil
	default:
		return nil, fmt.Errorf("invalid MedalType value: %v", i)
	}
}

type CommandNodeType string

const (
	CommandNodeTypeStart        CommandNodeType = "start"
	CommandNodeTypeConditionalA CommandNodeType = "conditional_a"
	CommandNodeTypeConditionalB CommandNodeType = "conditional_b"
	CommandNodeTypeConditionalC CommandNodeType = "conditional_c"
	CommandNodeTypeConditionalD CommandNodeType = "conditional_d"
	CommandNodeTypeConditionalE CommandNodeType = "conditional_e"
	CommandNodeTypeForward      CommandNodeType = "forward"
	CommandNodeTypeLeft         CommandNodeType = "left"
	CommandNodeTypeBack         CommandNodeType = "back"
	CommandNodeTypeRight        CommandNodeType = "right"
)

func (i *CommandNodeType) Scan(value interface{}) error {
	switch t := value.(type) {
	case []byte:
		*i = CommandNodeType(t)
	case string:
		*i = CommandNodeType(t)
	default:
		return fmt.Errorf("unsupported scan type for CommandNodeType: %T", t)
	}
	return nil
}

func (i CommandNodeType) Value() (driver.Value, error) {
	switch i {
	case CommandNodeTypeStart,
		CommandNodeTypeConditionalA,
		CommandNodeTypeConditionalB,
		CommandNodeTypeConditionalC,
		CommandNodeTypeConditionalD,
		CommandNodeTypeConditionalE,
		CommandNodeTypeForward,
		CommandNodeTypeLeft,
		CommandNodeTypeBack,
		CommandNodeTypeRight:
		return string(i), nil
	default:
		return nil, fmt.Errorf("invalid CommandNodeType value: %v", i)
	}
}

type CommandEdgeType string

const (
	CommandEdgeTypeConditionalBranch CommandEdgeType = "conditional_branch"
	CommandEdgeTypeMainBranch        CommandEdgeType = "main_branch"
)

func (i *CommandEdgeType) Scan(value interface{}) error {
	switch t := value.(type) {
	case []byte:
		*i = CommandEdgeType(t)
	case string:
		*i = CommandEdgeType(t)
	default:
		return fmt.Errorf("unsupported scan type for CommandEdgeType: %T", t)
	}
	return nil
}

func (i CommandEdgeType) Value() (driver.Value, error) {
	switch i {
	case CommandEdgeTypeConditionalBranch, CommandEdgeTypeMainBranch:
		return string(i), nil
	default:
		return nil, fmt.Errorf("invalid CommandEdgeType value: %v", i)
	}
}
