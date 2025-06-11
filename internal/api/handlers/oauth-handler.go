package handlers

import (
	"net/http"
	"rest-crud-go/internal/core/services"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type OAuthHandler struct {
	service *services.OAuthService
}

func CreateOAuthHandler(service *services.OAuthService) *OAuthHandler {
	return &OAuthHandler{service: service}
}

func (auth *OAuthHandler) GitHubLogin(ctx *gin.Context) {
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}

func (auth *OAuthHandler) GitHubCallback(ctx *gin.Context) {
	user, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.JSON(500, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	jwtToken, err := auth.service.AuthenticateGithub(ctx, user)
	if err != nil {
		ctx.JSON(500, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login successful",
		"token":   jwtToken,
		"user": gin.H{
			"id":       user.UserID,
			"email":    user.Email,
			"name":     user.Name,
			"avatar":   user.AvatarURL,
			"provider": "github",
		},
	})
}
