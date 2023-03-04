package mapper

import (
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
)

type GameSessionMapper struct{}

func NewSessionHistoryMapper() GameSessionMapper {
	return GameSessionMapper{}
}

func (m GameSessionMapper) ToDTO(sessionHistory *entity.GameSession) *dto.SessionHistoryDTO {
	submitHistoryMapper := NewSubmitHistoryMapper()

	submitHistoryDTOs := make([]*dto.SubmitHistoryDTO, 0, len(sessionHistory.SubmitHistories))
	for _, v := range sessionHistory.SubmitHistories {
		submitHistoryDTO := submitHistoryMapper.ToDTO(v)
		submitHistoryDTOs = append(submitHistoryDTOs, submitHistoryDTO)
	}

	sessionHistoryDTO := &dto.SessionHistoryDTO{
		MapConfigurationID: sessionHistory.MapConfigurationID,
		StartDatetime:      sessionHistory.StartDatetime,
		EndDatetime:        sessionHistory.EndDatetime,
		SubmitHistories:    submitHistoryDTOs,
	}

	return sessionHistoryDTO
}
