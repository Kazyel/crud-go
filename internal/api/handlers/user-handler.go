package handlers

import (
	"net/http"
	"rest-crud-go/internal/core/models"
	"rest-crud-go/internal/core/services"
	"rest-crud-go/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

type Pagination struct {
	Limit  int `form:"limit" binding:"omitempty,gt=0,lte=20"`
	Offset int `form:"offset" binding:"omitempty,gt=0,lte=100"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type PaginatedResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Meta    MetaData    `json:"meta"`
}

type MetaData struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

const (
	defaultLimit  = 10
	defaultOffset = 0
	maxLimit      = 100
)

func CreateUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var userRequest models.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		utils.HandleBindingError(ctx, err)
		return
	}

	userId, err := h.service.CreateUser(ctx, &userRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, APIResponse{
		Success: true,
		Message: "User created successfully!",
		Data: gin.H{
			"user_id": userId,
		},
	})
}

func (h *UserHandler) GetUserByID(ctx *gin.Context) {
	userId := ctx.Param("id")

	if userId == "" {
		ctx.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "User ID is required",
		})
		return
	}

	userData, err := h.service.GetUserByID(ctx, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   err.Error(),
		},
		)
		return
	}

	ctx.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data: gin.H{
			"user_data": userData,
		},
	})
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	var userRequest models.UserUpdateRequest
	userId := ctx.Param("id")

	if userId == "" {
		ctx.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "User ID is required",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		utils.HandleBindingError(ctx, err)
		return
	}

	updateUser, err := h.service.UpdateUser(ctx, userId, &userRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "User updated successfully",
		Data: gin.H{
			"user": updateUser,
		},
	})
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	if userId == "" {
		ctx.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "User ID is required",
		})
		return
	}

	if err := h.service.DeleteUser(ctx, userId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "User deleted successfully",
	})
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	var page Pagination

	if err := ctx.ShouldBindQuery(&page); err != nil {
		ctx.AbortWithStatusJSON(400, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	limit, offset := h.sanitizePagination(page)

	users, err := h.service.GetAllUsers(ctx, limit, offset)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, PaginatedResponse{
		Success: true,
		Data:    users,
		Meta: MetaData{
			Total:  len(users),
			Limit:  limit,
			Offset: offset,
		},
	})
}

func (h *UserHandler) sanitizePagination(page Pagination) (int, int) {
	limit := page.Limit
	offset := page.Offset

	if limit == 0 {
		limit = defaultLimit
	}
	if offset < 0 {
		offset = defaultOffset
	}

	if limit > maxLimit {
		limit = maxLimit
	}

	return limit, offset
}
