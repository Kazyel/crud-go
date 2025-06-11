package routes

import (
	"rest-crud-go/internal/api/handlers"
	"rest-crud-go/internal/api/middlewares"

	"github.com/gin-gonic/gin"
)

type AuthHandlers struct {
	AuthHandler  *handlers.AuthHandler
	OAuthHandler *handlers.OAuthHandler
}

func UserRoutes(router *gin.Engine, handler *handlers.UserHandler) {
	userRoute := router.Group("/api/v1/users")

	userRoute.GET("/:id", middlewares.AuthJWT(), handler.GetUserByID)
	userRoute.GET("/all", middlewares.AuthJWT(), handler.GetAllUsers)

	userRoute.POST("/", handler.CreateUser)

	userRoute.PATCH("/:id", middlewares.AuthJWT(), handler.UpdateUser)

	userRoute.DELETE("/:id", middlewares.AuthJWT(), handler.DeleteUser)
}

func AuthRoutes(router *gin.Engine, handlers *AuthHandlers) {
	authRoute := router.Group("/api/v1/auth")

	authRoute.GET("/github", handlers.OAuthHandler.GitHubLogin)
	authRoute.GET("/github/callback", handlers.OAuthHandler.GitHubCallback)

	authRoute.POST("/login", handlers.AuthHandler.Login)
	authRoute.POST("/logout", handlers.AuthHandler.Logout)
}
