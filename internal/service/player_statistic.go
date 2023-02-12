package service

import (
	"context"
	"llg_backend/internal/entity"
	"llg_backend/internal/repository"
)

type playerStatisticService struct {
	unitOfWork repository.UnitOfWork
}

func NewPlayerStatisticService(unitOfWork repository.UnitOfWork) PlayerStatisticService {
	return &playerStatisticService{
		unitOfWork: unitOfWork,
	}
}

func (s playerStatisticService) CreateSessionHistory(ctx context.Context, arg CreateSessionHistoryParams) (*entity.PlayerGameSession, error) {
	var resultGameSession *entity.PlayerGameSession

	err := s.unitOfWork.Do(ctx, func(store *repository.UnitOfWorkStore) error {
		gameSession, txErr := store.GameSessionRepo.CreateGameSession(ctx, repository.CreateGameSessionParams{
			PlayerID:           arg.PlayerID,
			MapConfigurationID: arg.MapConfigurationID,
			StartDatetime:      arg.StartDatetime,
			EndDatetime:        arg.EndDatetime,
		})
		if txErr != nil {
			return txErr
		}

		for _, v := range arg.GameHistories {
			playHistory, txErr := store.PlayHistoryRepo.CreatePlayHistory(ctx, repository.CreatePlayHistoryParams{
				GameSessionID:   gameSession.GameSesssionID,
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

			stateValue, txErr := store.PlayHistoryRepo.CreateStateValue(ctx, repository.CreateStateValueParams{
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
				ruleHistory, txErr := store.PlayHistoryRepo.CreateRuleHistory(ctx, repository.CreateRuleHistoryParams{
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
