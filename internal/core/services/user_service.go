package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"rest-crud-go/internal/core/models"
	"rest-crud-go/internal/core/repositories"
	"rest-crud-go/internal/utils"
	"strings"
	"time"
)

type UserService struct {
	repo repositories.UserRepository
}

type UserServiceInterface interface {
	CreateUser(ctx context.Context, req *models.UserRequest) (string, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, id string, req *models.UserUpdateRequest) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
	GetAllUsers(ctx context.Context, limit, offset int) ([]models.UsersData, error)
	AuthenticateUser(ctx context.Context, req *models.UserLoginRequest) (*models.User, error)
}

func CreateUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrUnauthorized = errors.New("unauthorized")
)

func (s *UserService) CreateUser(ctx context.Context, req *models.UserRequest) (string, error) {
	if s.userExistsByEmail(ctx, req.Email) {
		return "", utils.ErrUserExists
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return "", utils.ErrHashingFailed
	}

	user := &models.User{
		Name:      strings.TrimSpace(req.Name),
		Email:     strings.ToLower(strings.TrimSpace(req.Email)),
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		if errors.Is(err, utils.ErrUserExists) {
			return "", err
		}

		return "", fmt.Errorf("failed to create user: %w", err)
	}

	return user.ID, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, utils.ErrUserNotFound) {
			return nil, err
		}

		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id string, req *models.UserUpdateRequest) (*models.User, error) {
	if !s.userExistsByID(ctx, id) {
		return nil, utils.ErrUserNotFound
	}

	userUpdate := &models.UserUpdate{
		LastUpdated: time.Now(),
	}

	if req.Name != "" {
		name := strings.TrimSpace(req.Name)
		userUpdate.Name = &name
	}

	if req.Email != "" {
		email := strings.ToLower(strings.TrimSpace(req.Email))
		if s.isEmailTakenByOtherUser(ctx, email, id) {
			return nil, utils.ErrUserExists
		}
		userUpdate.Email = &email
	}

	if req.Password != "" {
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", utils.ErrHashingFailed, err)
		}
		userUpdate.Password = &hashedPassword
	}

	updatedUser, err := s.repo.UpdateUser(ctx, id, userUpdate)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	updatedUser.Password = sql.NullString{}
	return updatedUser, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	err := s.repo.DeleteUser(ctx, id)

	if err != nil {
		if errors.Is(err, utils.ErrUserNotFound) {
			return utils.ErrUserNotFound
		}

		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

func (s *UserService) GetAllUsers(ctx context.Context, limit, offset int) ([]models.UsersData, error) {
	users, err := s.repo.GetAllUsers(ctx, limit, offset)

	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	return users, nil
}

func (s *UserService) isEmailTakenByOtherUser(ctx context.Context, email, currentUserID string) bool {
	req := models.UserLoginRequest{Email: email}
	creds, err := s.repo.GetUserByEmail(ctx, req)
	if err != nil {
		return false
	}
	return creds.ID != currentUserID
}

func (s *UserService) userExistsByEmail(ctx context.Context, email string) bool {
	req := models.UserLoginRequest{Email: email}
	_, err := s.repo.GetUserByEmail(ctx, req)
	return err == nil
}

func (s *UserService) userExistsByID(ctx context.Context, id string) bool {
	_, err := s.repo.GetUserByID(ctx, id)
	return err == nil
}
