package repositories

import (
	"context"
	"fmt"
	"rest-crud-go/internal/core/models"
	"rest-crud-go/internal/core/utils"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/markbates/goth"
)

type OAuthRepositoryInterface interface {
	GetUserByProviderID(ctx context.Context, provider, providerID string) (*models.OAuthUser, error)
}

type OAuthRepository struct {
	db *pgxpool.Pool
}

func CreateOAuthRepository(db *pgxpool.Pool) *OAuthRepository {
	return &OAuthRepository{db: db}
}

func (r *OAuthRepository) GetUserByProviderID(ctx context.Context, provider, providerID string) (*models.OAuthUser, error) {
	var user models.OAuthUser

	query := `
			SELECT id, email, name, provider, provider_id, avatar_url, created_at
			FROM oauth_users 
			WHERE provider = $1 AND provider_id = $2
			RETURNING id, email, name, provider, provider_id, avatar_url, created_at
	`

	err := r.db.QueryRow(ctx, query, provider, providerID).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Provider,
		&user.ProviderID,
		&user.AvatarURL,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (h *OAuthRepository) CreateUser(ctx context.Context, gothUser goth.User) (*models.OAuthUser, error) {
	var user models.OAuthUser

	query := `
		INSERT INTO oauth_users (email, provider, provider_id, name, avatar_url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING email, provider, provider_id, name, avatar_url
	`

	err := h.db.QueryRow(ctx, query,
		gothUser.Email,
		"github",
		gothUser.UserID,
		gothUser.Name,
		gothUser.AvatarURL,
	).Scan(
		&user.Email,
		&user.Provider,
		&user.ProviderID,
		&user.Name,
		&user.AvatarURL,
	)

	if err != nil {
		if utils.IsPgError(err, utils.UniqueViolationErrCode) {
			return nil, utils.ErrUserExists
		}
		return nil, fmt.Errorf("error creating user: %s", err)
	}

	return &user, nil
}
