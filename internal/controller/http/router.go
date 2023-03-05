package http

import (
	v1 "llg_backend/internal/controller/http/v1"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(handler *gin.Engine, db *gorm.DB) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	serverStatusController := NewServerStatusController()

	handler.Static("/static", "./static")
	serverStatusController.initRoutes(&handler.RouterGroup)
	v1.NewRouter(&handler.RouterGroup, db)
}
