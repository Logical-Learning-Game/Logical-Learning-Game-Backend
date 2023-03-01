package v1

import (
	"fmt"
	"llg_backend/internal/controller/http/httputil"
	"llg_backend/internal/dto"
	"llg_backend/internal/dto/mapper"
	"llg_backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	mapConfigService service.MapConfigurationService
	statisticService service.PlayerStatisticService
}

func NewPlayerController(mapConfigService service.MapConfigurationService, statisticService service.PlayerStatisticService) *PlayerController {
	return &PlayerController{
		mapConfigService: mapConfigService,
		statisticService: statisticService,
	}
}

func (c PlayerController) initRoutes(handler *gin.RouterGroup) {
	h := handler.Group("/players")
	{
		playerGroup := h.Group("/:playerID")
		{
			playerGroup.GET("/available_maps", c.ListAvailableMaps)
			playerGroup.POST("/statistics", c.CreateSessionHistory)
		}
	}
}

func (c PlayerController) ListAvailableMaps(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	playerWorlds, err := c.mapConfigService.ListFromPlayerID(ctx, playerID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	worldMapper := mapper.NewWorldMapper()
	worldDTOs := make([]*dto.WorldDTO, 0, len(playerWorlds))
	for _, world := range playerWorlds {
		worldDTO := worldMapper.ToDTO(world)
		worldDTOs = append(worldDTOs, worldDTO)
	}

	ctx.JSON(http.StatusOK, worldDTOs)
}

func (c PlayerController) CreateSessionHistory(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	var createSessionHistoryRequestDTO dto.CreateGameSessionRequestDTO
	if err := ctx.ShouldBindJSON(&createSessionHistoryRequestDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httputil.ErrorResponse(err))
		return
	}

	gameSession, err := c.statisticService.CreateSessionHistory(ctx, playerID, createSessionHistoryRequestDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.Header("Location", fmt.Sprintf("/v1/sessions/%d", gameSession.ID))
	ctx.Status(http.StatusCreated)
}
