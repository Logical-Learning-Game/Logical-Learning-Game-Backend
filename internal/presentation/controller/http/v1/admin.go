package v1

import (
	"llg_backend/internal/dto"
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
		worldGroup := h.Group("/worlds")
		{
			worldGroup.GET("", c.ListWorld)
			worldGroup.POST("", c.CreateWorld)
			worldGroup.GET("/maps", c.ListWorldWithMap)
			singleWorldGroup := worldGroup.Group("/:worldID")
			{
				singleWorldGroup.PUT("", c.UpdateWorld)
			}
		}
		mapGroup := h.Group("/maps")
		{
			singleMapGroup := mapGroup.Group("/:mapID")
			{
				singleMapGroup.PATCH("/active", c.SetMapActive)
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

func (c AdminController) ListWorld(ctx *gin.Context) {
	worlds, err := c.mapConfigService.ListWorld(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, worlds)
}

func (c AdminController) CreateWorld(ctx *gin.Context) {
	var createWorldRequest dto.CreateWorldRequest
	if err := ctx.ShouldBindJSON(&createWorldRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httputil.ErrorResponse(err))
		return
	}

	if err := c.mapConfigService.CreateWorld(ctx, createWorldRequest.Name); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

func (c AdminController) UpdateWorld(ctx *gin.Context) {
	worldIDString := ctx.Param("worldID")
	worldID, err := strconv.ParseInt(worldIDString, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httputil.ErrorResponse(err))
		return
	}

	var updateWorldRequest dto.UpdateWorldRequest
	if err = ctx.ShouldBindJSON(&updateWorldRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httputil.ErrorResponse(err))
		return
	}

	if err = c.mapConfigService.UpdateWorld(ctx, worldID, updateWorldRequest.Name); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

func (c AdminController) ListWorldWithMap(ctx *gin.Context) {
	worldWithMaps, err := c.mapConfigService.ListWorldWithMap(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	for _, world := range worldWithMaps {
		for i := range world.Maps {
			if world.Maps[i].MapImagePath.Valid {
				absoluteImagePath := httputil.AbsoluteImageURL(ctx, world.Maps[i].MapImagePath.String)
				world.Maps[i].MapImagePath.String = absoluteImagePath
			}
		}
	}

	ctx.JSON(http.StatusOK, worldWithMaps)
}

func (c AdminController) SetMapActive(ctx *gin.Context) {
	mapIDString := ctx.Param("mapID")
	mapID, err := strconv.ParseInt(mapIDString, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httputil.ErrorResponse(err))
		return
	}

	var setMapActiveRequest dto.SetMapActiveRequest
	if err = ctx.ShouldBindJSON(&setMapActiveRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httputil.ErrorResponse(err))
		return
	}

	if err = c.mapConfigService.SetMapActive(ctx, mapID, setMapActiveRequest.Active); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}
