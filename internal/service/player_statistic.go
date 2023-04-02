package service

import (
	"context"
	"errors"
	"llg_backend/internal/dto"
	"llg_backend/internal/dto/mapper"
	"llg_backend/internal/entity"
	"sort"
	"time"

	"gorm.io/gorm"
)

type playerStatisticService struct {
	db *gorm.DB
}

func NewPlayerStatisticService(db *gorm.DB) PlayerStatisticService {
	return &playerStatisticService{db: db}
}

func (s playerStatisticService) CreateSessionHistory(ctx context.Context, playerID string, arg dto.SessionHistoryRequest) (*entity.GameSession, error) {
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
				AllItemCount:          int32(submitHistoryDTO.StateValue.AllItemCount),
				KeyACount:             int32(submitHistoryDTO.StateValue.KeyACount),
				KeyBCount:             int32(submitHistoryDTO.StateValue.KeyBCount),
				KeyCCount:             int32(submitHistoryDTO.StateValue.KeyCCount),
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
					Index:           int32(commandNodeDTO.Index),
					Type:            commandNodeDTO.Type,
					InGamePosition: entity.Vector2Float32{
						X: commandNodeDTO.PositionX,
						Y: commandNodeDTO.PositionY,
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

func (s playerStatisticService) UpdateTopSubmitHistory(ctx context.Context, playerID string, args []*dto.TopSubmitHistoryRequest) ([]*entity.SubmitHistory, error) {
	insertedTopSubmitHistories := make([]*entity.SubmitHistory, 0, len(args))

	txErr := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, entry := range args {
			topSubmit := entry.SubmitHistory

			// find map for player that need to update top submit history
			var mapConfigurationForPlayer entity.MapConfigurationForPlayer
			result := tx.Where(&entity.MapConfigurationForPlayer{
				PlayerID:           playerID,
				MapConfigurationID: entry.MapConfigurationID,
			}).First(&mapConfigurationForPlayer)
			if err := result.Error; err != nil {
				return err
			}

			// set is pass status to pass if top submit is completed
			if topSubmit.IsCompleted && !mapConfigurationForPlayer.IsPass {
				mapConfigurationForPlayer.IsPass = true
				result = tx.Save(&mapConfigurationForPlayer)
				if err := result.Error; err != nil {
					return err
				}
			}

			// query old submit history if not found then immediate
			var oldTopSubmitHistory entity.SubmitHistory
			foundOldTopSubmitHistory := true
			result = tx.Where(&entity.SubmitHistory{
				MapConfigurationForPlayerID: mapConfigurationForPlayer.ID,
			}).
				Joins("StateValue").
				Preload("SubmitHistoryRules").
				First(&oldTopSubmitHistory)
			if err := result.Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					foundOldTopSubmitHistory = false
				} else {
					return err
				}
			}

			if foundOldTopSubmitHistory {
				// compare old submit history with new submit history
				submitHistoryMapper := mapper.NewSubmitHistoryMapper()
				newSubmitHistoryToCompare := submitHistoryMapper.ToEntity(topSubmit)
				if !compareSubmitHistory(&oldTopSubmitHistory, newSubmitHistoryToCompare) {
					continue
				} else {
					// remove old top submit history
					result = tx.Delete(&entity.SubmitHistory{}, oldTopSubmitHistory.ID)
					if err := result.Error; err != nil {
						return err
					}
				}
			}

			// insert new top submit history
			newSubmitHistory := &entity.SubmitHistory{
				MapConfigurationForPlayerID: mapConfigurationForPlayer.ID,
				IsFinited:                   topSubmit.IsFinited,
				IsCompleted:                 topSubmit.IsCompleted,
				CommandMedal:                topSubmit.CommandMedal,
				ActionMedal:                 topSubmit.ActionMedal,
				SubmitDatetime:              topSubmit.SubmitDatetime,
			}

			result = tx.Omit("GameSessionID").Create(newSubmitHistory)
			if err := result.Error; err != nil {
				return err
			}

			// insert new state value
			topSubmitStateValue := topSubmit.StateValue
			stateValue := &entity.StateValue{
				SubmitHistoryID:       newSubmitHistory.ID,
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
				AllItemCount:          int32(topSubmitStateValue.AllItemCount),
				KeyACount:             int32(topSubmitStateValue.KeyACount),
				KeyBCount:             int32(topSubmitStateValue.KeyBCount),
				KeyCCount:             int32(topSubmitStateValue.KeyCCount),
			}

			result = tx.Create(stateValue)
			if err := result.Error; err != nil {
				return err
			}

			newSubmitHistory.StateValue = stateValue

			// insert new submit history rules
			for _, ruleDTO := range topSubmit.SubmitHistoryRules {
				rule := &entity.SubmitHistoryRule{
					SubmitHistoryID:        newSubmitHistory.ID,
					MapConfigurationRuleID: ruleDTO.MapRuleID,
					IsPass:                 ruleDTO.IsPass,
				}

				result = tx.Create(rule)
				if err := result.Error; err != nil {
					return err
				}

				newSubmitHistory.SubmitHistoryRules = append(newSubmitHistory.SubmitHistoryRules, rule)
			}

			// insert new command node
			for _, commandNodeDTO := range topSubmit.CommandNodes {
				commandNode := &entity.CommandNode{
					SubmitHistoryID: newSubmitHistory.ID,
					Index:           int32(commandNodeDTO.Index),
					Type:            commandNodeDTO.Type,
					InGamePosition: entity.Vector2Float32{
						X: commandNodeDTO.PositionX,
						Y: commandNodeDTO.PositionY,
					},
				}

				result = tx.Create(commandNode)
				if err := result.Error; err != nil {
					return err
				}

				newSubmitHistory.CommandNodes = append(newSubmitHistory.CommandNodes, commandNode)
			}

			// insert new command edge
			for _, commandEdgeDTO := range topSubmit.CommandEdges {
				commandEdge := &entity.CommandEdge{
					SubmitHistoryID:      newSubmitHistory.ID,
					SourceNodeIndex:      int32(commandEdgeDTO.SourceNodeIndex),
					DestinationNodeIndex: int32(commandEdgeDTO.DestinationNodeIndex),
					Type:                 commandEdgeDTO.Type,
				}

				result = tx.Create(commandEdge)
				if err := result.Error; err != nil {
					return err
				}

				newSubmitHistory.CommandEdges = append(newSubmitHistory.CommandEdges, commandEdge)
			}

			insertedTopSubmitHistories = append(insertedTopSubmitHistories, newSubmitHistory)
		}

		return nil
	})

	return insertedTopSubmitHistories, txErr
}

func (s playerStatisticService) ListTopSubmitHistory(ctx context.Context, playerID string) ([]*dto.TopSubmitHistoryResponse, error) {
	mapConfigurationForPlayers := make([]*entity.MapConfigurationForPlayer, 0)

	result := s.db.WithContext(ctx).
		Model(&entity.MapConfigurationForPlayer{}).
		InnerJoins("TopSubmitHistory").
		Where(&entity.MapConfigurationForPlayer{
			PlayerID: playerID,
		}).
		Preload("TopSubmitHistory.StateValue").
		Preload("TopSubmitHistory.SubmitHistoryRules").
		Preload("TopSubmitHistory.SubmitHistoryRules.MapConfigurationRule").
		Preload("TopSubmitHistory.CommandNodes").
		Preload("TopSubmitHistory.CommandEdges").
		Find(&mapConfigurationForPlayers)
	if err := result.Error; err != nil {
		return nil, err
	}

	submitHistoryMapper := mapper.NewSubmitHistoryMapper()

	topSubmitHistoryDTOs := make([]*dto.TopSubmitHistoryResponse, 0, len(mapConfigurationForPlayers))
	for _, v := range mapConfigurationForPlayers {
		topSubmitHistoryDTOs = append(topSubmitHistoryDTOs, &dto.TopSubmitHistoryResponse{
			MapConfigurationID: v.MapConfigurationID,
			SubmitHistory:      submitHistoryMapper.ToSubmitHistoryResponse(v.TopSubmitHistory),
		})
	}

	return topSubmitHistoryDTOs, nil
}

func (s playerStatisticService) ListPlayerSessionDataForGame(ctx context.Context, playerID string) ([]*dto.SessionHistoryForGameResponse, error) {
	gameSessions := make([]*entity.GameSession, 0)

	sixMonthsAgo := time.Now().AddDate(0, -6, 0)
	result := s.db.WithContext(ctx).
		Where("player_id = ? AND start_datetime >= ?", playerID, sixMonthsAgo).
		Order("start_datetime DESC").
		Preload("SubmitHistories").
		Preload("SubmitHistories.StateValue").
		Preload("SubmitHistories.SubmitHistoryRules").
		Preload("SubmitHistories.SubmitHistoryRules.MapConfigurationRule").
		Preload("SubmitHistories.CommandNodes").
		Preload("SubmitHistories.CommandEdges").
		Find(&gameSessions)
	if err := result.Error; err != nil {
		return nil, err
	}

	sessionHistoryMapper := mapper.NewSessionHistoryMapper()

	sessionHistoryDTOs := make([]*dto.SessionHistoryForGameResponse, 0, len(gameSessions))
	for _, v := range gameSessions {
		sessionHistoryDTO := sessionHistoryMapper.ToDTO(v)
		sessionHistoryDTOs = append(sessionHistoryDTOs, sessionHistoryDTO)
	}

	return sessionHistoryDTOs, nil
}

func (s playerStatisticService) ListPlayerSessionForAdmin(ctx context.Context, playerID string) ([]*dto.SessionDataForAdminResponse, error) {
	gameSessions := make([]*entity.GameSession, 0)

	result := s.db.WithContext(ctx).
		Joins("MapConfiguration").
		Where(&entity.GameSession{
			PlayerID: playerID,
		}).
		Order("start_datetime DESC").
		Preload("MapConfiguration.World").
		Preload("SubmitHistories", func(db *gorm.DB) *gorm.DB {
			return db.Order("submit_datetime DESC")
		}).
		Preload("SubmitHistories.StateValue").
		Preload("SubmitHistories.SubmitHistoryRules").
		Preload("SubmitHistories.SubmitHistoryRules.MapConfigurationRule").
		Preload("SubmitHistories.CommandNodes").
		Preload("SubmitHistories.CommandEdges").
		Find(&gameSessions)
	if err := result.Error; err != nil {
		return nil, err
	}

	submitHistoryMapper := mapper.NewSubmitHistoryMapper()
	sessionDataForAdmins := make([]*dto.SessionDataForAdminResponse, 0, len(gameSessions))
	for _, session := range gameSessions {
		submitHistoryForAdmins := make([]*dto.SubmitHistoryForAdminResponse, 0, len(session.SubmitHistories))

		for _, submit := range session.SubmitHistories {
			submitForAdmin := submitHistoryMapper.ToSubmitHistoryForAdminResponse(submit)
			submitHistoryForAdmins = append(submitHistoryForAdmins, submitForAdmin)
		}

		sessionDataForAdmins = append(sessionDataForAdmins, &dto.SessionDataForAdminResponse{
			SessionID:       session.ID,
			WorldID:         session.MapConfiguration.WorldID,
			WorldName:       session.MapConfiguration.World.Name,
			MapID:           session.MapConfigurationID,
			MapName:         session.MapConfiguration.ConfigName,
			StartDatetime:   session.StartDatetime,
			EndDatetime:     session.EndDatetime,
			SubmitHistories: submitHistoryForAdmins,
		})
	}

	return sessionDataForAdmins, nil
}

func (s playerStatisticService) ListSubmitHistoriesForAdmin(ctx context.Context, sessionID int64) ([]*dto.SubmitHistoryForAdminResponse, error) {
	submitHistories := make([]*entity.SubmitHistory, 0)

	result := s.db.WithContext(ctx).
		Joins("StateValue").
		Where(&entity.SubmitHistory{
			GameSessionID: sessionID,
		}).
		Preload("SubmitHistoryRules").
		Preload("SubmitHistoryRules.MapConfigurationRule").
		Preload("CommandNodes").
		Preload("CommandEdges").
		Find(&submitHistories)
	if err := result.Error; err != nil {
		return nil, err
	}

	submitHistoryMapper := mapper.NewSubmitHistoryMapper()

	submitHistoryForAdmins := make([]*dto.SubmitHistoryForAdminResponse, 0, len(submitHistories))
	for _, v := range submitHistories {
		submitForAdmin := submitHistoryMapper.ToSubmitHistoryForAdminResponse(v)
		submitHistoryForAdmins = append(submitHistoryForAdmins, submitForAdmin)
	}

	return submitHistoryForAdmins, nil
}

func (s playerStatisticService) ListMapOfPlayerInfoForAdmin(ctx context.Context, playerID string) ([]*dto.MapOfPlayerInfoForAdminResponse, error) {
	mapConfigurationForPlayers := make([]*entity.MapConfigurationForPlayer, 0)

	result := s.db.WithContext(ctx).
		Joins("MapConfiguration").
		Joins("TopSubmitHistory").
		Where(&entity.MapConfigurationForPlayer{
			PlayerID: playerID,
		}).
		Preload("MapConfiguration.World").
		Preload("TopSubmitHistory.StateValue").
		Preload("TopSubmitHistory.SubmitHistoryRules").
		Preload("TopSubmitHistory.SubmitHistoryRules.MapConfigurationRule").
		Preload("TopSubmitHistory.CommandNodes").
		Preload("TopSubmitHistory.CommandEdges").
		Find(&mapConfigurationForPlayers)
	if err := result.Error; err != nil {
		return nil, err
	}

	// sort by world's id and map's id field
	sort.SliceStable(mapConfigurationForPlayers, func(i, j int) bool {
		if mapConfigurationForPlayers[i].MapConfiguration.WorldID != mapConfigurationForPlayers[j].MapConfiguration.WorldID {
			return mapConfigurationForPlayers[i].MapConfiguration.WorldID < mapConfigurationForPlayers[j].MapConfiguration.WorldID
		}
		return mapConfigurationForPlayers[i].MapConfigurationID < mapConfigurationForPlayers[j].MapConfigurationID
	})

	submitHistoryMapper := mapper.NewSubmitHistoryMapper()

	mapOfPlayerInfoResponse := make([]*dto.MapOfPlayerInfoForAdminResponse, 0, len(mapConfigurationForPlayers))
	for _, v := range mapConfigurationForPlayers {
		var submit *dto.SubmitHistoryForAdminResponse
		if v.TopSubmitHistory != nil {
			submit = submitHistoryMapper.ToSubmitHistoryForAdminResponse(v.TopSubmitHistory)
		}

		m := &dto.MapOfPlayerInfoForAdminResponse{
			MapForPlayerID:   v.ID,
			WorldID:          v.MapConfiguration.WorldID,
			WorldName:        v.MapConfiguration.World.Name,
			MapID:            v.MapConfigurationID,
			MapName:          v.MapConfiguration.ConfigName,
			IsPass:           v.IsPass,
			TopSubmitHistory: submit,
		}

		mapOfPlayerInfoResponse = append(mapOfPlayerInfoResponse, m)
	}

	return mapOfPlayerInfoResponse, nil
}

func (s playerStatisticService) GetPlayerData(ctx context.Context, playerID string) (*dto.PlayerDataDTO, error) {
	sessionHistoryDTOs, err := s.ListPlayerSessionDataForGame(ctx, playerID)
	if err != nil {
		return nil, err
	}

	topSubmitHistoryDTOs, err := s.ListTopSubmitHistory(ctx, playerID)
	if err != nil {
		return nil, err
	}

	syncPlayerDataDTO := &dto.PlayerDataDTO{
		SessionHistories:   sessionHistoryDTOs,
		TopSubmitHistories: topSubmitHistoryDTOs,
	}

	return syncPlayerDataDTO, nil
}

// compareSubmitHistory will return true if newSubmitHistory is better than oldSubmitHistory otherwise false
func compareSubmitHistory(oldSubmitHistory, newSubmitHistory *entity.SubmitHistory) bool {
	if newSubmitHistory.IsCompleted && !oldSubmitHistory.IsCompleted {
		return true
	} else if !newSubmitHistory.IsCompleted && oldSubmitHistory.IsCompleted {
		return false
	} else {
		// compare rule histories
		oldRulePass := 0
		newRulePass := 0
		for i := 0; i < len(newSubmitHistory.SubmitHistoryRules); i++ {
			if newSubmitHistory.SubmitHistoryRules[i].IsPass {
				newRulePass++
			}

			if oldSubmitHistory.SubmitHistoryRules[i].IsPass {
				oldRulePass++
			}
		}

		if newRulePass > oldRulePass {
			return true
		} else if newRulePass < oldRulePass {
			return false
		} else {
			// compare command medal
			if entity.MedalValue[newSubmitHistory.CommandMedal] > entity.MedalValue[oldSubmitHistory.CommandMedal] {
				return true
			} else if entity.MedalValue[newSubmitHistory.CommandMedal] < entity.MedalValue[oldSubmitHistory.CommandMedal] {
				return false
			} else {
				// compare action medal
				if entity.MedalValue[newSubmitHistory.ActionMedal] > entity.MedalValue[oldSubmitHistory.ActionMedal] {
					return true
				} else if entity.MedalValue[newSubmitHistory.ActionMedal] < entity.MedalValue[oldSubmitHistory.ActionMedal] {
					return false
				} else {
					// compare state value
					if newSubmitHistory.StateValue.CommandCount < oldSubmitHistory.StateValue.CommandCount {
						return true
					} else if newSubmitHistory.StateValue.CommandCount > oldSubmitHistory.StateValue.CommandCount {
						return false
					} else {
						return newSubmitHistory.StateValue.ActionCount < oldSubmitHistory.StateValue.ActionCount
					}
				}
			}
		}
	}
}
