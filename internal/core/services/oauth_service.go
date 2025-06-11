package services

import (
	"context"
	"fmt"
	"rest-crud-go/internal/core/models"
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
	user, err := s.createOrFindUser(ctx, request)
	if err != nil {
		return "", err
	}

	jwtToken, _, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (s *OAuthService) createOrFindUser(ctx context.Context, request goth.User) (*models.OAuthUser, error) {
	// TODO: Implement user update later
	existingUser, err := s.repo.GetUserByProviderID(ctx, "github", request.UserID)

	if err == nil {
		return existingUser, nil
	}

	user, err := s.repo.CreateUser(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

// func (s *OAuthService) userExistsByEmail(ctx context.Context, email string) bool {
// 	req := models.UserLoginRequest{Email: email}
// 	_, err := s.repo.GetUserByEmail(ctx, req)
// 	return err == nil
// }
