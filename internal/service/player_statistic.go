package service

import (
	"context"
	"llg_backend/internal/entity"
)

type playerStatisticService struct {
	unitOfWork entity.UnitOfWork
}

func NewPlayerStatisticService(unitOfWork entity.UnitOfWork) entity.PlayerStatisticService {
	return &playerStatisticService{
		unitOfWork: unitOfWork,
	}
}

func (s playerStatisticService) CreateSessionHistory(ctx context.Context, arg entity.CreateSessionHistoryParams) (*entity.GameSession, error) {
	var resultGameSession *entity.GameSession

	err := s.unitOfWork.Do(ctx, func(store *entity.UnitOfWorkStore) error {
		gameSession, txErr := store.GameSessionRepo.CreateGameSession(ctx, entity.CreateGameSessionParams{
			PlayerID:           arg.PlayerID,
			MapConfigurationID: arg.MapConfigurationID,
			StartDatetime:      arg.StartDatetime,
			EndDatetime:        arg.EndDatetime,
		})
		if txErr != nil {
			return txErr
		}

		for _, history := range arg.GameHistories {
			playHistory, txErr := store.PlayHistoryRepo.CreatePlayHistory(ctx, entity.CreatePlayHistoryParams{
				GameSessionID:   gameSession.ID,
				ActionStep:      history.ActionStep,
				NumberOfCommand: history.NumberOfCommand,
				IsFinited:       history.IsFinited,
				IsCompleted:     history.IsCompleted,
				CommandMedal:    history.CommandMedal,
				ActionMedal:     history.ActionMedal,
				SubmitDatetime:  history.SubmitDatetime,
			})
			if txErr != nil {
				return txErr
			}

			stateValue, txErr := store.PlayHistoryRepo.CreateStateValue(ctx, entity.CreateStateValueParams{
				PlayHistoryID:         playHistory.ID,
				CommandCount:          history.StateValue.CommandCount,
				ForwardCommandCount:   history.StateValue.ForwardCommandCount,
				RightCommandCount:     history.StateValue.RightCommandCount,
				BackCommandCount:      history.StateValue.BackCommandCount,
				LeftCommandCount:      history.StateValue.LeftCommandCount,
				ConditionCommandCount: history.StateValue.ConditionCommandCount,
				ActionCount:           history.StateValue.ActionCount,
				ForwardActionCount:    history.StateValue.ForwardActionCount,
				RightActionCount:      history.StateValue.RightActionCount,
				BackActionCount:       history.StateValue.BackActionCount,
				LeftActionCount:       history.StateValue.LeftActionCount,
				ConditionActionCount:  history.StateValue.ConditionActionCount,
			})
			if txErr != nil {
				return txErr
			}

			playHistory.StateValue = stateValue

			for _, rule := range history.Rules {
				ruleHistory, txErr := store.PlayHistoryRepo.CreateRuleHistory(ctx, entity.CreateRuleHistoryParams{
					PlayHistoryID:   playHistory.ID,
					MapConfigRuleID: rule.MapConfigRuleID,
					IsPass:          rule.IsPass,
				})
				if txErr != nil {
					return txErr
				}

				playHistory.Rules = append(playHistory.Rules, ruleHistory)
			}

			nodeIndexMap := make(map[int]int64)
			for _, node := range history.CommandNodes {
				commandNode, txErr := store.PlayHistoryRepo.CreateCommandNode(ctx, entity.CreateCommandNodeParams{
					PlayHistoryID: playHistory.ID,
					Type:          node.Type,
					InGamePosition: entity.Vector2Float{
						X: node.InGamePosition.X,
						Y: node.InGamePosition.Y,
					},
				})
				if txErr != nil {
					return txErr
				}

				nodeIndexMap[node.NodeIndex] = commandNode.ID
				playHistory.CommandNodes = append(playHistory.CommandNodes, commandNode)
			}

			for _, edge := range history.CommandEdges {
				commandEdge, txErr := store.PlayHistoryRepo.CreateCommandEdge(ctx, entity.CreateCommandEdgeParams{
					SourceNodeID:      nodeIndexMap[edge.SourceNodeIndex],
					DestinationNodeID: nodeIndexMap[edge.DestinationIndex],
					Type:              edge.Type,
				})
				if txErr != nil {
					return txErr
				}

				playHistory.CommandEdges = append(playHistory.CommandEdges, commandEdge)
			}

			gameSession.GameHistory = append(gameSession.GameHistory, playHistory)
		}

		resultGameSession = gameSession
		return nil
	})

	return resultGameSession, err
}
