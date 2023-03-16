package v1

import (
	"llg_backend/internal/presentation/controller/http/httputil"
	"llg_backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	mapConfigService service.MapConfigurationService
	statisticService service.PlayerStatisticService
	playerService    service.PlayerService
}

func NewAdminController(mapConfigService service.MapConfigurationService, statisticService service.PlayerStatisticService, playerService service.PlayerService) *AdminController {
	return &AdminController{
		mapConfigService: mapConfigService,
		statisticService: statisticService,
		playerService:    playerService,
	}
}

func (c AdminController) initRoutes(handler *gin.RouterGroup) {
	h := handler.Group("/admin")
	{
		multiplePlayerGroup := h.Group("/players")
		{
			multiplePlayerGroup.GET("", c.ListPlayers)
			singlePlayerGroup := multiplePlayerGroup.Group("/:playerID")
			{
				singlePlayerMapGroup := singlePlayerGroup.Group("/map")
				{
					singlePlayerMapGroup.GET("/info", c.ListMapOfPlayerInfo)
				}
				singlePlayerGroup.GET("/sessions", c.ListPlayerSessions)
			}
		}
		sessionGroup := h.Group("/sessions")
		{
			singleSessionGroup := sessionGroup.Group("/:sessionID")
			{
				singleSessionGroup.GET("/submit_histories", c.ListSubmitHistories)
			}
		}
	}
}

func (c AdminController) ListPlayers(ctx *gin.Context) {
	playerInfoResponse, err := c.playerService.ListPlayers(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, playerInfoResponse)
}

func (c AdminController) ListPlayerSessions(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	playerSession, err := c.statisticService.ListPlayerSessionForAdmin(ctx, playerID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, playerSession)
}

func (c AdminController) ListSubmitHistories(ctx *gin.Context) {
	sessionIDString := ctx.Param("sessionID")
	sessionID, err := strconv.ParseInt(sessionIDString, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httputil.ErrorResponse(err))
		return
	}

	submitHistories, err := c.statisticService.ListSubmitHistoriesForAdmin(ctx, sessionID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, submitHistories)
}

func (c AdminController) ListMapOfPlayerInfo(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	mapOfPlayerInfo, err := c.statisticService.ListMapOfPlayerInfoForAdmin(ctx, playerID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, mapOfPlayerInfo)
}
