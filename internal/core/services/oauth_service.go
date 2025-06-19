package services

import (
	"context"
	"fmt"
	"rest-crud-go/internal/core/repositories"
	"rest-crud-go/internal/utils"

	"github.com/markbates/goth"
)

type OAuthService struct {
	repo *repositories.OAuthRepository
}

func CreateOAuthService(repo *repositories.OAuthRepository) *OAuthService {
	return &OAuthService{repo: repo}
}

func (s *OAuthService) AuthenticateGithub(ctx context.Context, request goth.User) (string, error) {
	userId, err := s.handleUserCredentials(ctx, request)
	if err != nil {
		return "", err
	}

	jwtToken, _, err := utils.GenerateJWT(userId)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (s *OAuthService) handleUserCredentials(ctx context.Context, request goth.User) (string, error) {
	existingUserId, _ := s.repo.GetUserByProviderID(ctx, "github", request.UserID)

	if existingUserId != "" {
		err := s.repo.UpdateUser(ctx, request)
		if err != nil {
			return "", err
		}

		return existingUserId, nil
	}

	newUserId, err := s.repo.CreateUser(ctx, request)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	return newUserId, nil
}
