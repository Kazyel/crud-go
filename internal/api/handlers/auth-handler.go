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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	userTokens, err := auth.service.Login(ctx, loginRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
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

	ctx.JSON(http.StatusOK, gin.H{
		"message":    "login successful",
		"csrf_token": userTokens.CSRFToken,
		"user_id":    userTokens.UserID,
	})
}

func (auth *AuthHandler) Logout(ctx *gin.Context) {
	ctx.SetCookie(
		"auth_token",
		"",
		-1,
		"/",
		"",
		true,
		true,
	)

	ctx.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
