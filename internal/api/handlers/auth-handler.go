package handlers

import (
	"net/http"
	"rest-crud-go/internal/core/models"
	"rest-crud-go/internal/core/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.AuthService
}

func CreateAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (auth *AuthHandler) Login(ctx *gin.Context) {
	var loginRequest models.UserLoginRequest

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	userTokens, err := auth.service.AuthenticateUser(ctx, loginRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	ctx.SetCookie(
		"jwt",
		userTokens.JWTToken,
		86400,
		"/",
		"",
		true,
		true,
	)

	ctx.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Login successful",
		Data: gin.H{
			"csrf_token": userTokens.CSRFToken,
			"user_id":    userTokens.UserID,
		}})
}

func (auth *AuthHandler) Logout(ctx *gin.Context) {
	ctx.SetCookie(
		"jwt",
		"",
		-1,
		"/",
		"",
		true,
		true,
	)

	ctx.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Logged out successfully",
	})
}
