package handlers

import (
	"net/http"
	"rest-crud-go/internal/core/models"
	"rest-crud-go/internal/core/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var user models.CreateUserRequest

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	_, err = h.service.CreateUser(ctx, &user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully!",
	})
}

func (h *UserHandler) GetUserByID(ctx *gin.Context) {
	// userId := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{})

}
