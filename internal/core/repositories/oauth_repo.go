package repositories

import (
	"context"
	"errors"
	"fmt"
	"rest-crud-go/internal/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/markbates/goth"
)

type OAuthRepositoryInterface interface {
	GetUserByProviderID(ctx context.Context, provider, providerID string) (string, error)
	CreateUser(ctx context.Context, gothUser goth.User) (string, error)
	UpdateUser(ctx context.Context, gothRequest goth.User) error
}

type OAuthRepository struct {
	db *pgxpool.Pool
}

func CreateOAuthRepository(db *pgxpool.Pool) *OAuthRepository {
	return &OAuthRepository{db: db}
}

func (h *OAuthRepository) CreateUser(ctx context.Context, gothUser goth.User) (string, error) {
	var userId string

	query := `
		INSERT INTO oauth_users (email, provider, provider_id, name, avatar_url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING provider_id
	`

	err := h.db.QueryRow(ctx, query,
		gothUser.Email,
		"github",
		gothUser.UserID,
		gothUser.Name,
		gothUser.AvatarURL,
	).Scan(
		&userId,
	)

	if err != nil {
		if utils.IsPgError(err, utils.UniqueViolationErrCode) {
			return "", utils.ErrUserExists
		}
		return "", fmt.Errorf("error creating user: %s", err)
	}

	return userId, nil
}

func (r *OAuthRepository) GetUserByProviderID(ctx context.Context, provider, providerID string) (string, error) {
	query := `
			SELECT id
			FROM oauth_users 
			WHERE provider = $1 AND provider_id = $2
	`

	var userId string
	err := r.db.QueryRow(ctx, query, provider, providerID).Scan(&userId)

	if err != nil {
		return "", err
	}

	return userId, nil
}

func (r *OAuthRepository) UpdateUser(ctx context.Context, gothRequest goth.User) error {
	query := `
		UPDATE oauth_users SET
		name = COALESCE($1, name),
		avatar_url = COALESCE($2, avatar_url),
		email = COALESCE($3, email)
		WHERE provider = $4 AND provider_id = $5		 
	`
	_, err := r.db.Exec(ctx, query,
		gothRequest.Name,
		gothRequest.AvatarURL,
		gothRequest.Email,
		"github",
		gothRequest.UserID,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.ErrUserNotFound
		}
		return fmt.Errorf("error updating user: %s", err)
	}

	return nil
}
