package middlewares

import (
	"net/http"
	"rest-crud-go/internal/core/utils"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenCookie, err := ctx.Cookie("jwt")

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing authentication token",
			})
			return
		}

		claims, err := utils.ParseJWT(tokenCookie)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalind token",
			})
			return
		}

		csrfToken := ctx.GetHeader("X-CSRF-Token")

		if csrfToken == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Missing CSRF token",
			})
			return
		}

		ctx.Set("userId", claims.UserID)
		ctx.Set("csrfToken", claims.CSRFToken)
		ctx.Next()
	}
}
