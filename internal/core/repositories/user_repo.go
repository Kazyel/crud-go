package repositories

import (
	"context"
	"fmt"
	"rest-crud-go/internal/core/models"

	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	CreateUser(context context.Context, user *models.User) error
	GetByID(context context.Context, id string) (*models.User, error)

	// GetAll(context context.Context) ([]models.User, error)
	// UpdateUser(context context.Context, user *models.User) error
	// DeleteUser(context context.Context, id string) error
}

type userRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `
				INSERT INTO users (name, email, password_hash, created_at) 
				VALUES ($1, $2, $3, $4)
				`

	_, err := r.db.Exec(ctx, query, user.Name, user.Email, user.Password, user.CreatedAt)

	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	query := `
		SELECT id, name, email, created_at 
		FROM users 
		WHERE id = $1		
	`

	user := models.User{}

	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user with ID %s not found", id)
		}

		return &models.User{}, fmt.Errorf("error querying user by id: %w", err)
	}

	return &user, nil
}
