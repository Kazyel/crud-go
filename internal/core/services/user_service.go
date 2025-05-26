package services

import (
	"context"
	"fmt"
	"rest-crud-go/internal/core/models"
	"rest-crud-go/internal/core/repositories"
	"rest-crud-go/internal/core/utils"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	repo repositories.UserRepository
}

func CreateUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req *models.UserRequest) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		return nil, fmt.Errorf("hashing failed: %w", err)
	}

	user := &models.User{
		Name:      req.Name,
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
	user, err := s.repo.GetUserByID(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("repository failed: %w", err)
	}

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id string, req *models.UserUpdateRequest) (*models.User, error) {
	userUpdate := models.UserUpdate{
		LastUpdated: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if req.Name != "" {
		userUpdate.Name = pgtype.Text{String: req.Name, Valid: true}
	}

	if req.Email != "" {
		userUpdate.Email = pgtype.Text{String: req.Email, Valid: true}
	}

	if req.Password != "" {
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return nil, fmt.Errorf("hashing failed: %w", err)
		}
		userUpdate.Password = pgtype.Text{String: hashedPassword, Valid: true}
	}

	updatedUser, err := s.repo.UpdateUser(ctx, id, &userUpdate)

	if err != nil {
		return nil, fmt.Errorf("repository failed: %w", err)
	}

	return updatedUser, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	err := s.repo.DeleteUser(ctx, id)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return fmt.Errorf("user with ID %s not found", id)
		}

		return fmt.Errorf("repository failed: %w", err)
	}

	return nil
}

func (s *UserService) GetAllUsers(ctx context.Context, limit, offset int) ([]models.UsersData, error) {
	users, err := s.repo.GetAllUsers(ctx, limit, offset)

	if err != nil {
		return nil, fmt.Errorf("repository failed: %w", err)
	}

	return users, nil
}
