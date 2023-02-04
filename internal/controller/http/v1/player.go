package v1

import (
	"llg_backend/internal/pkg/httputil"
	"llg_backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	playerService service.PlayerService
	worldService  service.WorldService
}

func NewPlayerController(playerService service.PlayerService, worldService service.WorldService) *PlayerController {
	return &PlayerController{
		playerService: playerService,
		worldService:  worldService,
	}
}

func (c *PlayerController) initRoutes(handler *gin.RouterGroup) {
	h := handler.Group("/player")
	{
		h.POST("/login_log", c.CreateLoginLog)
		playerGroup := h.Group("/:playerID")
		{
			playerGroup.GET("/available_maps", c.ListAvailableMaps)
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
	}

	ctx.JSON(http.StatusOK, playerWorlds)
}
