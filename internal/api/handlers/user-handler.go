package handlers

import (
	"net/http"
	"rest-crud-go/internal/core/models"
	"rest-crud-go/internal/core/services"
	"rest-crud-go/internal/core/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func CreateUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var userRequest models.UserRequest

	err := ctx.ShouldBindJSON(&userRequest)

	if err != nil {
		utils.HandleBindingError(ctx, err)
		return
	}

	_, err = h.service.CreateUser(ctx, &userRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully!",
	})
}

func (h *UserHandler) GetUserByID(ctx *gin.Context) {
	userId := ctx.Param("id")

	user, err := h.service.GetUserByID(ctx, userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	var userRequest models.UserUpdateRequest
	userId := ctx.Param("id")

	err := ctx.ShouldBindJSON(&userRequest)

	if err != nil {
		utils.HandleBindingError(ctx, err)
		return
	}

	user, err := h.service.UpdateUser(ctx, userId, &userRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
