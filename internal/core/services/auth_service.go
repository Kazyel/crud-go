package services

import (
	"context"
	"errors"
	"fmt"
	"rest-crud-go/internal/core/models"
	"rest-crud-go/internal/core/repositories"
	"rest-crud-go/internal/utils"
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

func (s *AuthService) AuthenticateUser(ctx context.Context, request models.UserLoginRequest) (*UserTokens, error) {
	creds, err := s.repo.GetUserByEmail(ctx, request)
	if err != nil {
		if errors.Is(err, utils.ErrUserNotFound) {
			return nil, ErrUnauthorized
		}
		return nil, fmt.Errorf("failed to authenticate user: %w", err)
	}

	if !utils.VerifyPassword(request.Password, creds.Password) {
		return nil, utils.ErrHashingFailed
	}

	user, err := s.repo.GetUserByID(ctx, creds.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user data: %w", err)
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
