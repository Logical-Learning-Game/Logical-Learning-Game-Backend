package v1

import (
	"github.com/gin-gonic/gin"
	"llg_backend/config"
	"llg_backend/internal/service"
)

func NewRouter(handler *gin.Engine, cfg *config.Config, playerService service.PlayerService, worldService service.WorldService) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/v1")
	{
		newPlayerController(h, playerService, worldService)
	}
}
