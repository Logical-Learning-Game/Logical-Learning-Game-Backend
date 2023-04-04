package mapper

import (
	"llg_backend/internal/dto"
	"llg_backend/internal/entity"
)

type WorldMapper struct{}

func NewWorldMapper() WorldMapper {
	return WorldMapper{}
}

func (m WorldMapper) ToWorldDTO(world *entity.World) *dto.WorldDTO {
	if world == nil {
		return nil
	}

	mapConfigDTOs := make([]*dto.MapConfigurationDTO, 0, len(world.MapConfigurations))
	mapConfigMapper := NewMapConfigurationMapper()
	for _, mapConfig := range world.MapConfigurations {
		mapConfigDTO := mapConfigMapper.ToMapConfigurationDTO(mapConfig)
		mapConfigDTOs = append(mapConfigDTOs, mapConfigDTO)
	}

	worldDTO := &dto.WorldDTO{
		WorldID:   world.ID,
		WorldName: world.Name,
		Maps:      mapConfigDTOs,
	}

	return worldDTO
}

func (m WorldMapper) ToWorldWithMapForAdminResponse(world *entity.World) *dto.WorldWithMapForAdminResponse {
	if world == nil {
		return nil
	}

	mapConfigMapper := NewMapConfigurationMapper()

	mapConfigForAdminDTOs := make([]*dto.MapConfigurationForAdminDTO, 0, len(world.MapConfigurations))
	for _, mapConfig := range world.MapConfigurations {
		mapConfigForAdminDTO := mapConfigMapper.ToMapConfigurationForAdminDTO(mapConfig)
		mapConfigForAdminDTOs = append(mapConfigForAdminDTOs, mapConfigForAdminDTO)
	}

	worldWithMapForAdmin := &dto.WorldWithMapForAdminResponse{
		WorldID:   world.ID,
		WorldName: world.Name,
		Maps:      mapConfigForAdminDTOs,
	}

	return worldWithMapForAdmin
}
