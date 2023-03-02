package service

import (
	"context"
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
	"time"

	"gorm.io/gorm"
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
				GameSessionID:  gameSession.ID,
				IsFinited:      submitHistoryDTO.IsFinited,
				IsCompleted:    submitHistoryDTO.IsCompleted,
				CommandMedal:   submitHistoryDTO.CommandMedal,
				ActionMedal:    submitHistoryDTO.ActionMedal,
				SubmitDatetime: submitHistoryDTO.SubmitDatetime,
			}

			result = tx.Omit("MapConfigurationForPlayerID").Create(submitHistory)
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

func (s playerStatisticService) UpdateTopSubmitHistory(ctx context.Context, playerID string, args []*dto.TopSubmitHistoryDTO) ([]*entity.SubmitHistory, error) {
	insertedTopSubmitHistories := make([]*entity.SubmitHistory, 0, len(args))

	txErr := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, entry := range args {
			// find map for player that need to update top submit history
			var mapConfigurationForPlayer entity.MapConfigurationForPlayer
			tx.Where(&entity.MapConfigurationForPlayer{
				PlayerID:           playerID,
				MapConfigurationID: entry.MapConfigurationID,
			}).Find(&mapConfigurationForPlayer)

			// set is pass status to pass
			mapConfigurationForPlayer.IsPass = true
			tx.Save(&mapConfigurationForPlayer)

			// remove old top submit history
			tx.Where(&entity.SubmitHistory{
				MapConfigurationForPlayerID: mapConfigurationForPlayer.ID,
			}).Delete(&entity.SubmitHistory{})

			// insert new top submit history
			topSubmit := entry.SubmitHistory
			submitHistory := &entity.SubmitHistory{
				MapConfigurationForPlayerID: mapConfigurationForPlayer.ID,
				IsFinited:                   topSubmit.IsFinited,
				IsCompleted:                 topSubmit.IsCompleted,
				CommandMedal:                topSubmit.CommandMedal,
				ActionMedal:                 topSubmit.ActionMedal,
				SubmitDatetime:              topSubmit.SubmitDatetime,
			}

			result := tx.Omit("GameSessionID").Create(submitHistory)
			if err := result.Error; err != nil {
				return err
			}

			// insert new state value
			topSubmitStateValue := topSubmit.StateValue
			stateValue := &entity.StateValue{
				SubmitHistoryID:       submitHistory.ID,
				CommandCount:          int32(topSubmitStateValue.CommandCount),
				ForwardCommandCount:   int32(topSubmitStateValue.ForwardCommandCount),
				RightCommandCount:     int32(topSubmitStateValue.RightCommandCount),
				BackCommandCount:      int32(topSubmitStateValue.BackCommandCount),
				LeftCommandCount:      int32(topSubmitStateValue.LeftCommandCount),
				ConditionCommandCount: int32(topSubmitStateValue.ConditionCommandCount),
				ActionCount:           int32(topSubmitStateValue.ActionCount),
				ForwardActionCount:    int32(topSubmitStateValue.ForwardActionCount),
				RightActionCount:      int32(topSubmitStateValue.RightActionCount),
				BackActionCount:       int32(topSubmitStateValue.BackActionCount),
				LeftActionCount:       int32(topSubmitStateValue.LeftActionCount),
				ConditionActionCount:  int32(topSubmitStateValue.ConditionActionCount),
			}

			result = tx.Create(stateValue)
			if err := result.Error; err != nil {
				return err
			}

			submitHistory.StateValue = stateValue

			// insert new submit history rules
			for _, ruleDTO := range topSubmit.SubmitHistoryRules {
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

			// insert new command node
			for _, commandNodeDTO := range topSubmit.CommandNodes {
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

			// insert new command edge
			for _, commandEdgeDTO := range topSubmit.CommandEdges {
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

			insertedTopSubmitHistories = append(insertedTopSubmitHistories, submitHistory)
		}

		return nil
	})

	return insertedTopSubmitHistories, txErr
}

func (s playerStatisticService) ListPlayerSessionData(ctx context.Context, playerID string) ([]*entity.GameSession, error) {
	gameSessions := make([]*entity.GameSession, 0)

	sixMonthsAgo := time.Now().AddDate(0, -6, 0)
	result := s.db.WithContext(ctx).
		Preload("SubmitHistories").
		Preload("SubmitHistories.StateValue").
		Preload("SubmitHistories.SubmitHistoryRules").
		Preload("SubmitHistories.CommandNodes").
		Preload("SubmitHistories.CommandEdges").
		Where("player_id = ? AND start_datetime >= ?", playerID, sixMonthsAgo).
		Find(&gameSessions)

	return gameSessions, result.Error
}
