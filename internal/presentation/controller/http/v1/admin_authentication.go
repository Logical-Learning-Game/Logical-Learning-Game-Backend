package v1

import (
	"errors"
	"llg_backend/internal/dto"
	"llg_backend/internal/presentation/controller/http/httputil"
	"llg_backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminAuthenticationController struct {
	authenticationService service.AdminAuthenticationService
}

func NewAdminAuthenticationController(authenticationService service.AdminAuthenticationService) *AdminAuthenticationController {
	return &AdminAuthenticationController{authenticationService: authenticationService}
}

func (c AdminAuthenticationController) initRoutes(handler *gin.RouterGroup) {
	h := handler.Group("/auth/admin")
	{
		h.POST("/login", c.Login)
	}
}

func (c AdminAuthenticationController) Login(ctx *gin.Context) {
	var loginRequest dto.AdminLoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httputil.ErrorResponse(err))
		return
	}

	adminLoginResponse, err := c.authenticationService.Login(ctx, loginRequest.Username, loginRequest.Password)
	if err != nil {
		if errors.Is(err, service.ErrAdminNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, httputil.ErrorResponse(err))
			return
		} else if errors.Is(err, service.ErrUnauthorized) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, httputil.ErrorResponse(err))
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputil.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, adminLoginResponse)
}
