package v1

import (
	"llg_backend/internal/pkg/httputil"
	"llg_backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type playerController struct {
	playerService service.PlayerService
}

func newPlayerController(handler *gin.RouterGroup, playerService service.PlayerService) {
	controller := playerController{
		playerService: playerService,
	}

	h := handler.Group("/player")
	{
		h.POST("/login_log", controller.CreateLoginLog)
	}
}

type createLoginLogRequest struct {
	PlayerID string `json:"player_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func (c *playerController) CreateLoginLog(ctx *gin.Context) {
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
