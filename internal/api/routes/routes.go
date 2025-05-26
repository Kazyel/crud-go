package routes

import (
	"rest-crud-go/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, handler *handlers.UserHandler) {
	userRoute := router.Group("/api/v1/users")

	userRoute.GET("/:id", handler.GetUserByID)
	userRoute.GET("/all", handler.GetAllUsers)

	userRoute.POST("/", handler.CreateUser)
	userRoute.POST("/login", handler.UserLogin)

	userRoute.PATCH("/:id", handler.UpdateUser)

	userRoute.DELETE("/:id", handler.DeleteUser)
}
