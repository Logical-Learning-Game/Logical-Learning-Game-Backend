package http

import (
	"llg_backend/config"
	"llg_backend/internal/presentation/controller/http/middleware"
	"llg_backend/internal/presentation/controller/http/v1"
	"llg_backend/internal/token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(handler *gin.Engine, cfg *config.Config, db *gorm.DB, tokenMaker token.Maker) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(middleware.CORSMiddleware())

	serverStatusController := NewServerStatusController()

	handler.Static("/static", "./static")
	serverStatusController.initRoutes(&handler.RouterGroup)
	v1.NewRouter(&handler.RouterGroup, cfg, db, tokenMaker)
}
