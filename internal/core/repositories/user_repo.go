package repositories

import (
	"context"
	"errors"
	"fmt"
	"rest-crud-go/internal/core/models"
	"rest-crud-go/internal/core/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, userRequest models.UserLoginRequest) (userCredentials, error)
	GetAllUsers(ctx context.Context, limit, offset int) ([]models.UsersData, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, id string, userRequest *models.UserUpdate) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type userRepository struct {
	db *pgxpool.Pool
}

type userCredentials struct {
	ID       string `db:"id"`
	Password string `db:"password_hash"`
}

func CreateUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, userRequest *models.User) error {
	query := `
		INSERT INTO users (name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id`

	var id string
	err := r.db.QueryRow(ctx, query,
		userRequest.Name,
		userRequest.Email,
		userRequest.Password,
	).Scan(&id)

	if err != nil {
		if utils.IsPgError(err, utils.UniqueViolationErrCode) {
			return utils.ErrUserExists
		}
		return fmt.Errorf("error creating user: %w", err)
	}

	userRequest.ID = id
	return nil
}

func (r *userRepository) GetAllUsers(ctx context.Context, limit, offset int) ([]models.UsersData, error) {
	query := `
		SELECT id, name, email, created_at 
		FROM users 
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error getting all users: %w", err)
	}

	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.UsersData])
	if err != nil {
		return nil, fmt.Errorf("error collecting user rows: %w", err)
	}

	return users, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	query := `
		SELECT id, name, email, created_at, last_updated
		FROM users 
		WHERE id = $1		
	`

	user := &models.User{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.LastUpdated,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.ErrUserNotFound
		}
		return &models.User{}, fmt.Errorf("error querying user by id: %w", err)
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, userRequest models.UserLoginRequest) (userCredentials, error) {
	query := `
		SELECT id, password_hash
		FROM users 
		WHERE email = $1		
	`

	var creds userCredentials
	err := r.db.QueryRow(ctx, query, userRequest.Email).Scan(
		&creds.ID,
		&creds.Password,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return userCredentials{}, utils.ErrUserNotFound
		}
		return userCredentials{}, fmt.Errorf("error querying user by email: %w", err)
	}

	return creds, nil
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.ErrUserNotFound
		}
		if utils.IsPgError(err, utils.UniqueViolationErrCode) {
			return nil, utils.ErrUserExists
		}
		return nil, fmt.Errorf("error updating user: %w", err)
	}

	return &userReturn, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return utils.ErrUserNotFound
	}

	return nil
}
