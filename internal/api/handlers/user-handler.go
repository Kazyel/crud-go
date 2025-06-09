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

type Pagination struct {
	Limit  int `form:"limit" binding:"omitempty,gt=0,lte=20"`
	Offset int `form:"offset" binding:"omitempty,gt=0,lte=100"`
}

func CreateUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var userRequest models.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		utils.HandleBindingError(ctx, err)
		return
	}

	_, err := h.service.CreateUser(ctx, &userRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		utils.HandleBindingError(ctx, err)
	}

	user, err := h.service.UpdateUser(ctx, userId, &userRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	if err := h.service.DeleteUser(ctx, userId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully!",
	})
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	var page Pagination
	var LIMIT int = 15
	var OFFSET int = 0

	if err := ctx.ShouldBindQuery(&page); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	if page.Limit != 0 {
		LIMIT = page.Limit
	}

	if page.Offset != 0 {
		OFFSET = page.Offset
	}

	users, err := h.service.GetAllUsers(ctx, LIMIT, OFFSET)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": users,
	})
}
