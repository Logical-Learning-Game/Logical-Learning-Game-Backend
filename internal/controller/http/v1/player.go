package v1

import (
	"llg_backend/internal/entity"
	"llg_backend/internal/pkg/httputil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	playerService          entity.PlayerService
	worldService           entity.WorldService
	playerStatisticService entity.PlayerStatisticService
}

func NewPlayerController(playerService entity.PlayerService, worldService entity.WorldService, playerStatisticService entity.PlayerStatisticService) *PlayerController {
	return &PlayerController{
		playerService:          playerService,
		worldService:           worldService,
		playerStatisticService: playerStatisticService,
	}
}

func (c *PlayerController) initRoutes(handler *gin.RouterGroup) {
	h := handler.Group("/players")
	{
		h.POST("/login_log", c.CreateLoginLog)
		playerGroup := h.Group("/:playerID")
		{
			playerGroup.GET("/available_maps", c.ListAvailableMaps)
			playerGroup.POST("/statistics", c.CreateSessionHistory)
		}
	}
}

type createLoginLogRequest struct {
	PlayerID string `json:"player_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func (c *PlayerController) CreateLoginLog(ctx *gin.Context) {
	var req createLoginLogRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httputil.ErrorResponse(err))
		return
	}

	if err := c.playerService.CreateOrUpdatePlayerInformation(ctx, req.PlayerID, req.Email, req.Name); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	if err := c.playerService.CreateLoginLog(ctx, req.PlayerID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c *PlayerController) ListAvailableMaps(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	playerWorlds, err := c.worldService.ListFromPlayerID(ctx, playerID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, playerWorlds)
}

func (c *PlayerController) CreateSessionHistory(ctx *gin.Context) {
	playerID := ctx.Param("playerID")

	var req entity.CreateSessionHistoryParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httputil.ErrorResponse(err))
		return
	}

	req.PlayerID = playerID

	gameSession, err := c.playerStatisticService.CreateSessionHistory(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, gameSession)
}
