package services

import (
	"context"
	"fmt"
	"rest-crud-go/internal/core/models"
	"rest-crud-go/internal/core/repositories"
	"rest-crud-go/internal/core/utils"
)

type AuthService struct {
	repo repositories.UserRepository
}

func CreateAuthService(repo repositories.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(ctx context.Context, request models.UserLoginRequest) (string, string, error) {
	user, err := s.repo.GetUserByEmail(ctx, request)

	if err != nil {
		return "", "", fmt.Errorf("repository failed: %w", err)
	}

	ok, err := utils.VerifyPassword(request.Password, user.Password)

	if err != nil {
		return "", "", fmt.Errorf("error verifying password: %w", err)
	}

	if !ok {
		return "", "", fmt.Errorf("invalid password")
	}

	newToken, err := utils.GenerateJWT(user.ID)

	if err != nil {
		return "", "", err
	}

	return user.ID, newToken, nil
}
