package presenter

import "llg_backend/internal/dto"

type PlayerDataPresenter struct{}

func NewPlayerDataPresenter() PlayerDataPresenter {
	return PlayerDataPresenter{}
}

func (p PlayerDataPresenter) Present(playerID string, playerDataDTO *dto.PlayerDataDTO) *dto.PlayerDataResponse {
	sessionHistoryWithStatusDTOs := make([]*dto.SessionHistoryWithStatusResponse, 0, len(playerDataDTO.SessionHistories))
	for _, v := range playerDataDTO.SessionHistories {
		sessionHistoryWithStatusDTOs = append(sessionHistoryWithStatusDTOs, &dto.SessionHistoryWithStatusResponse{
			SessionHistory: v,
			Status:         true,
		})
	}

	topSubmitMap := make(map[int64]*dto.SubmitHistoryResponse)
	for _, v := range playerDataDTO.TopSubmitHistories {
		topSubmitMap[v.MapConfigurationID] = v.SubmitHistory
	}

	playerDataResponseDTO := &dto.PlayerDataResponse{
		PlayerID:           playerID,
		SessionHistories:   sessionHistoryWithStatusDTOs,
		TopSubmitHistories: topSubmitMap,
	}

	return playerDataResponseDTO
}
