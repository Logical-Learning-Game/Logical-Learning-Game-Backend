package v1

import (
	"database/sql"
	"llg_backend/internal/service"
	"llg_backend/internal/service/repository/player"
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
}

func (c *playerController) CreateLoginLog(ctx *gin.Context) {
	var req createLoginLogRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := player.CreateOrUpdatePlayerParams{
		PlayerID: req.PlayerID,
		Email: sql.NullString{
			String: req.Email,
			Valid:  req.Email != "",
		},
	}
	if err := c.playerService.CreateOrUpdatePlayer(ctx, arg); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err := c.playerService.CreateLoginLog(ctx, req.PlayerID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusCreated)
}
