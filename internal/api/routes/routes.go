package routes

import (
	"rest-crud-go/internal/api/handlers"
	"rest-crud-go/internal/api/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, handler *handlers.UserHandler) {
	userRoute := router.Group("/api/v1/users")
	userRoute.Use(middlewares.AuthTokenMiddleware())

	userRoute.GET("/:id", handler.GetUserByID)
	userRoute.GET("/all", handler.GetAllUsers)

	userRoute.POST("/", handler.CreateUser)

	userRoute.PATCH("/:id", handler.UpdateUser)

	userRoute.DELETE("/:id", handler.DeleteUser)
}

func AuthRoutes(router *gin.Engine, handler *handlers.AuthHandler) {
	authRoute := router.Group("/api/v1/auth")

	authRoute.POST("/login", handler.Login)
}
