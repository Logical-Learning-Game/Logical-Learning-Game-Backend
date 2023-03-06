package v1

import (
	"errors"
	"fmt"
	"llg_backend/internal/dto"
	"llg_backend/internal/presentation/controller/http/httputil"
	"llg_backend/internal/presentation/presenter"
	"llg_backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	mapConfigService service.MapConfigurationService
	statisticService service.PlayerStatisticService
	playerService    service.PlayerService
}

func NewPlayerController(mapConfigService service.MapConfigurationService, statisticService service.PlayerStatisticService, playerService service.PlayerService) *PlayerController {
	return &PlayerController{
		mapConfigService: mapConfigService,
		statisticService: statisticService,
		playerService:    playerService,
	}
}

func (c PlayerController) initRoutes(handler *gin.RouterGroup) {
	h := handler.Group("/players")
	{
		h.POST("/link_account", c.LinkAccount)
		playerGroup := h.Group("/:playerID")
		{
			playerGroup.GET("/", c.PlayerInfo)
			playerGroup.GET("/game_data", c.GetPlayerData)
			playerGroup.DELETE("/game_data", c.RemovePlayerData)
			playerGroup.GET("/session_history", c.ListSessionHistory)
			playerGroup.POST("/session_history", c.CreateSessionHistory)
			playerGroup.GET("/top_submit_history", c.ListTopSubmitHistory)
			playerGroup.POST("/top_submit_history", c.UpdateTopSubmitHistory)
			playerGroup.GET("/maps", c.ListAvailableMaps)
		}
	}
}

func (c PlayerController) ListAvailableMaps(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	worldDTOs, err := c.mapConfigService.ListPlayerAvailableMaps(ctx, playerID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	for _, world := range worldDTOs {
		for i := range world.Maps {
			if world.Maps[i].MapImagePath.Valid {
				absoluteImagePath := httputil.AbsoluteImageURL(ctx, world.Maps[i].MapImagePath.String)
				world.Maps[i].MapImagePath.String = absoluteImagePath
			}
		}
	}

	ctx.JSON(http.StatusOK, worldDTOs)
}

func (c PlayerController) CreateSessionHistory(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	var createSessionHistoryRequestDTO dto.SessionHistoryRequest
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

func (c PlayerController) ListSessionHistory(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	sessionHistoryDTOs, err := c.statisticService.ListPlayerSessionData(ctx, playerID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sessionHistoryDTOs)
}

func (c PlayerController) UpdateTopSubmitHistory(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	var topSubmitHistoryDTO []*dto.TopSubmitHistoryRequest
	if err := ctx.ShouldBindJSON(&topSubmitHistoryDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httputil.ErrorResponse(err))
		return
	}

	_, err := c.statisticService.UpdateTopSubmitHistory(ctx, playerID, topSubmitHistoryDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c PlayerController) ListTopSubmitHistory(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	topSubmitHistoryDTOs, err := c.statisticService.ListTopSubmitHistory(ctx, playerID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, topSubmitHistoryDTOs)
}

func (c PlayerController) GetPlayerData(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	playerDataDTO, err := c.statisticService.GetPlayerData(ctx, playerID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	playerDataPresenter := presenter.NewPlayerDataPresenter()
	formattedPlayerData := playerDataPresenter.Present(playerID, playerDataDTO)

	ctx.JSON(http.StatusOK, formattedPlayerData)
}

func (c PlayerController) RemovePlayerData(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	if err := c.playerService.RemovePlayerData(ctx, playerID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c PlayerController) LinkAccount(ctx *gin.Context) {
	var linkAccountRequestDTO dto.LinkAccountRequest
	if err := ctx.ShouldBindJSON(&linkAccountRequestDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httputil.ErrorResponse(err))
		return
	}

	user, err := c.playerService.LinkAccount(ctx, linkAccountRequestDTO)
	if err != nil {
		if errors.Is(err, service.ErrAccountAlreadyLinked) {
			ctx.AbortWithStatusJSON(http.StatusConflict, httputil.ErrorResponse(err))
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.Header("Location", fmt.Sprintf("/v1/players/%s", user.PlayerID))
	ctx.Status(http.StatusCreated)
}

func (c PlayerController) PlayerInfo(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	playerInfoResponse, err := c.playerService.PlayerInfo(ctx, playerID)
	if err != nil {
		if errors.Is(err, service.ErrPlayerNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, httputil.ErrorResponse(err))
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, playerInfoResponse)
}
