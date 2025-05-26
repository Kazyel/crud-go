package repositories

import (
	"context"
	"fmt"
	"rest-crud-go/internal/core/models"

	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	GetUserByID(context context.Context, id string) (*models.User, error)
	GetUserByEmail(context context.Context, email string) (string, error)
	GetAllUsers(context context.Context, limit, offset int) ([]models.UsersData, error)
	CreateUser(context context.Context, user *models.User) error
	UpdateUser(context context.Context, id string, user *models.UserUpdate) (*models.User, error)
	DeleteUser(context context.Context, id string) error
}

type userRepository struct {
	db *pgx.Conn
}

func CreateUserRepository(db *pgx.Conn) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `
				INSERT INTO users (name, email, password_hash, created_at) 
				VALUES ($1, $2, $3, $4)
				`

	_, err := r.db.Exec(ctx, query, user.Name, user.Email, user.Password, user.LastUpdated)

	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (r *userRepository) GetAllUsers(ctx context.Context, limit, offset int) ([]models.UsersData, error) {
	query := `
	SELECT id, name, email FROM users
	LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(ctx, query, limit, offset)

	if err != nil {
		return nil, fmt.Errorf("error getting all users: %w", err)
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.UsersData])
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
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

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (string, error) {
	query := `
		SELECT id
		FROM users 
		WHERE email = $1		
	`

	var userId string

	err := r.db.QueryRow(ctx, query, email).Scan(
		&userId,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("user with email %s not found", email)
		}

		return "", fmt.Errorf("error querying user by email: %w", err)
	}

	return userId, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id string, user *models.UserUpdate) (*models.User, error) {
	query := `
		UPDATE users SET
		name = COALESCE($1, name),
		email = COALESCE($2, email),
		password_hash = COALESCE($3, password_hash),
		last_updated = $4
		WHERE id = $5
		RETURNING id, name, email, created_at, last_updated
	`

	userReturn := models.User{}

	err := r.db.QueryRow(ctx, query,
		user.Name,
		user.Email,
		user.Password,
		user.LastUpdated,
		id).Scan(
		&userReturn.ID,
		&userReturn.Name,
		&userReturn.Email,
		&userReturn.CreatedAt,
		&userReturn.LastUpdated,
	)

	if err != nil {
		return nil, fmt.Errorf("error updating user: %w", err)
	}

	return &userReturn, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx,
		query,
		id)

	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("user with ID %s not found", id)
	}

	return nil
}
