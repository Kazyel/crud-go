package services

import (
	"context"
	"fmt"
	"rest-crud-go/internal/core/models"
	"rest-crud-go/internal/core/repositories"
	"rest-crud-go/internal/core/utils"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	repo repositories.UserRepository
}

func CreateUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		return nil, fmt.Errorf("hashing failed: %w", err)
	}

	user := &models.User{
		Name:      uuid.NewString(),
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	err = s.repo.CreateUser(ctx, user)

	if err != nil {
		return nil, fmt.Errorf("repository failed: %w", err)
	}

	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user, err := s.repo.GetByID(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("repository failed: %w", err)
	}

	return user, nil
}
