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

		for _, v := range arg.GameHistories {
			playHistory, txErr := store.PlayHistoryRepo.CreatePlayHistory(ctx, entity.CreatePlayHistoryParams{
				GameSessionID:   gameSession.ID,
				ActionStep:      v.ActionStep,
				NumberOfCommand: v.NumberOfCommand,
				IsFinited:       v.IsFinited,
				IsCompleted:     v.IsCompleted,
				CommandMedal:    v.CommandMedal,
				ActionMedal:     v.ActionMedal,
				SubmitDatetime:  v.SubmitDatetime,
			})
			if txErr != nil {
				return txErr
			}

			stateValue, txErr := store.PlayHistoryRepo.CreateStateValue(ctx, entity.CreateStateValueParams{
				PlayHistoryID:         playHistory.ID,
				CommandCount:          v.StateValue.CommandCount,
				ForwardCommandCount:   v.StateValue.ForwardCommandCount,
				RightCommandCount:     v.StateValue.RightCommandCount,
				BackCommandCount:      v.StateValue.BackCommandCount,
				LeftCommandCount:      v.StateValue.LeftCommandCount,
				ConditionCommandCount: v.StateValue.ConditionCommandCount,
				ActionCount:           v.StateValue.ActionCount,
				ForwardActionCount:    v.StateValue.ForwardActionCount,
				RightActionCount:      v.StateValue.RightActionCount,
				BackActionCount:       v.StateValue.BackActionCount,
				LeftActionCount:       v.StateValue.LeftActionCount,
				ConditionActionCount:  v.StateValue.ConditionActionCount,
			})
			if txErr != nil {
				return txErr
			}

			playHistory.StateValue = stateValue

			for _, rule := range v.Rules {
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

			gameSession.GameHistory = append(gameSession.GameHistory, playHistory)
		}

		resultGameSession = gameSession
		return nil
	})

	return resultGameSession, err
}
