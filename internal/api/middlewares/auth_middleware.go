package middlewares

import (
	"net/http"
	"rest-crud-go/internal/core/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthTokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized access",
				"message": "Missing or invalid bearer token, please provide a valid token"})
			return
		}

		bearerToken := strings.TrimPrefix(authHeader, "Bearer ")

		if bearerToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Invalid token",
				"message": "The provided token is empty",
			})
			return
		}

		userId, err := utils.ParseJWT(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Invalid token",
				"message": "The provided token is invalid",
			})
			return
		}

		ctx.Set("authToken", bearerToken)
		ctx.Set("userId", userId)
		ctx.Next()
	}
}
