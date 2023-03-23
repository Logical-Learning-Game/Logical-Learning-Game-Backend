package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(handler *gin.RouterGroup, db *gorm.DB) {
	playerController := InitializePlayerController(db)
	adminController := InitializeAdminController(db)

	h := handler.Group("/v1")
	{
		playerController.initRoutes(h)
		adminController.initRoutes(h)
	}
}
