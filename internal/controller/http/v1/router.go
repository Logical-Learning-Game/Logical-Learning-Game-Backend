package v1

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, playerCtrl *PlayerController) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/v1")
	{
		playerCtrl.initRoutes(h)
	}
}
