package httputil

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AbsoluteImageURL(ctx *gin.Context, relativePath string) string {
	return fmt.Sprintf("http://%s%s", ctx.Request.Host, relativePath)
}
