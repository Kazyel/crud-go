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

type UserTokens struct {
	JWTToken  string
	CSRFToken string
	UserID    string
}

func CreateAuthService(repo repositories.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(ctx context.Context, request models.UserLoginRequest) (*UserTokens, error) {
	user, err := s.repo.GetUserByEmail(ctx, request)

	if err != nil {
		return nil, fmt.Errorf("repository failed: %w", err)
	}

	ok, err := utils.VerifyPassword(request.Password, user.Password)

	if err != nil {
		return nil, fmt.Errorf("error verifying password: %w", err)
	}

	if !ok {
		return nil, fmt.Errorf("invalid password")
	}

	jwtToken, csrfToken, err := utils.GenerateJWT(user.ID)

	if err != nil {
		return nil, err
	}

	userTokens := UserTokens{
		JWTToken:  jwtToken,
		CSRFToken: csrfToken,
		UserID:    user.ID,
	}

	return &userTokens, nil
}
