package service

import (
	"context"
	"gorm.io/gorm"
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
)

type playerStatisticService struct {
	db *gorm.DB
}

func NewPlayerStatisticService(db *gorm.DB) PlayerStatisticService {
	return &playerStatisticService{db: db}
}

func (s playerStatisticService) CreateSessionHistory(ctx context.Context, playerID string, arg dto.CreateGameSessionRequestDTO) (*entity.GameSession, error) {
	gameSession := &entity.GameSession{
		PlayerID:           playerID,
		MapConfigurationID: arg.MapConfigurationID,
		StartDatetime:      arg.StartDatetime,
		EndDatetime:        arg.EndDatetime,
	}

	txErr := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Create(gameSession)
		if err := result.Error; err != nil {
			return err
		}

		for _, submitHistoryDTO := range arg.SubmitHistories {
			submitHistory := &entity.SubmitHistory{
				GameSessionID:   gameSession.ID,
				ActionStep:      int32(submitHistoryDTO.ActionStep),
				NumberOfCommand: int32(submitHistoryDTO.NumberOfCommand),
				IsFinited:       submitHistoryDTO.IsFinited,
				IsCompleted:     submitHistoryDTO.IsCompleted,
				CommandMedal:    submitHistoryDTO.CommandMedal,
				ActionMedal:     submitHistoryDTO.ActionMedal,
				SubmitDatetime:  submitHistoryDTO.SubmitDatetime,
			}

			result = tx.Create(submitHistory)
			if err := result.Error; err != nil {
				return err
			}

			stateValue := &entity.StateValue{
				SubmitHistoryID:       submitHistory.ID,
				CommandCount:          int32(submitHistoryDTO.StateValue.CommandCount),
				ForwardCommandCount:   int32(submitHistoryDTO.StateValue.ForwardCommandCount),
				RightCommandCount:     int32(submitHistoryDTO.StateValue.RightCommandCount),
				BackCommandCount:      int32(submitHistoryDTO.StateValue.BackCommandCount),
				LeftCommandCount:      int32(submitHistoryDTO.StateValue.LeftCommandCount),
				ConditionCommandCount: int32(submitHistoryDTO.StateValue.ConditionCommandCount),
				ActionCount:           int32(submitHistoryDTO.StateValue.ActionCount),
				ForwardActionCount:    int32(submitHistoryDTO.StateValue.ForwardActionCount),
				RightActionCount:      int32(submitHistoryDTO.StateValue.RightActionCount),
				BackActionCount:       int32(submitHistoryDTO.StateValue.BackActionCount),
				LeftActionCount:       int32(submitHistoryDTO.StateValue.LeftActionCount),
				ConditionActionCount:  int32(submitHistoryDTO.StateValue.ConditionActionCount),
			}

			result = tx.Create(stateValue)
			if err := result.Error; err != nil {
				return err
			}

			submitHistory.StateValue = stateValue

			for _, ruleDTO := range submitHistoryDTO.SubmitHistoryRules {
				rule := &entity.SubmitHistoryRule{
					SubmitHistoryID:        submitHistory.ID,
					MapConfigurationRuleID: ruleDTO.MapRuleID,
					IsPass:                 ruleDTO.IsPass,
				}

				result = tx.Create(rule)
				if err := result.Error; err != nil {
					return err
				}

				submitHistory.SubmitHistoryRules = append(submitHistory.SubmitHistoryRules, rule)
			}

			for _, commandNodeDTO := range submitHistoryDTO.CommandNodes {
				commandNode := &entity.CommandNode{
					SubmitHistoryID: submitHistory.ID,
					Index:           int32(commandNodeDTO.NodeIndex),
					Type:            commandNodeDTO.Type,
					InGamePosition: entity.Vector2Float32{
						X: commandNodeDTO.InGamePosition.X,
						Y: commandNodeDTO.InGamePosition.Y,
					},
				}

				result = tx.Create(commandNode)
				if err := result.Error; err != nil {
					return err
				}

				submitHistory.CommandNodes = append(submitHistory.CommandNodes, commandNode)
			}

			for _, commandEdgeDTO := range submitHistoryDTO.CommandEdges {
				commandEdge := &entity.CommandEdge{
					SubmitHistoryID:      submitHistory.ID,
					SourceNodeIndex:      int32(commandEdgeDTO.SourceNodeIndex),
					DestinationNodeIndex: int32(commandEdgeDTO.DestinationNodeIndex),
					Type:                 commandEdgeDTO.Type,
				}

				result = tx.Create(commandEdge)
				if err := result.Error; err != nil {
					return err
				}

				submitHistory.CommandEdges = append(submitHistory.CommandEdges, commandEdge)
			}

			gameSession.SubmitHistories = append(gameSession.SubmitHistories, submitHistory)
		}

		return nil
	})

	return gameSession, txErr
}
