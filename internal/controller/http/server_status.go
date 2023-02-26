package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerStatusController struct{}

func (c ServerStatusController) initRoutes(handler *gin.RouterGroup) {
	handler.GET("/status", c.Status)
}

func NewServerStatusController() *ServerStatusController {
	return &ServerStatusController{}
}

func (c ServerStatusController) Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "The server is up and running",
	})
}
