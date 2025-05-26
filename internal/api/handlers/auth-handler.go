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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, token, err := auth.service.Login(ctx, loginRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":    id,
		"token":   token,
		"message": "logged in successfully!",
	})
}
