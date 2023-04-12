package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"llg_backend/config"
	"llg_backend/internal/token"
)

func NewRouter(handler *gin.RouterGroup, cfg *config.Config, db *gorm.DB, tokenMaker token.Maker) {
	playerController := InitializePlayerController(db)
	adminController := InitializeAdminController(db)
	adminAuthController := InitializeAdminAuthenticationController(cfg, db, tokenMaker)

	h := handler.Group("/v1")
	{
		playerController.initRoutes(h)
		adminController.initRoutes(h, tokenMaker)
		adminAuthController.initRoutes(h)
	}
}
