package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID             string         `json:"id" db:"id"`
	Name           string         `json:"name" db:"name"`
	Email          string         `json:"email" db:"email"`
	Password       sql.NullString `json:"-" db:"password_hash"`
	Provider       string         `json:"provider" db:"provider"`
	ProviderID     sql.NullString `json:"provider_id" db:"provider_id"`
	AvatarURL      string         `json:"avatar_url" db:"avatar_url"`
	GithubUsername string         `json:"github_username" db:"github_username"`
	CreatedAt      time.Time      `json:"created_at" db:"created_at"`
	LastUpdated    time.Time      `json:"last_updated" db:"last_updated"`
}

type UserUpdate struct {
	Name           *string         `json:"name,omitempty"`
	Email          *string         `json:"email,omitempty"`
	Password       *sql.NullString `json:"-"`
	AvatarURL      sql.NullString  `json:"avatar_url" db:"avatar_url"`
	GithubUsername sql.NullString  `json:"github_username" db:"github_username"`
	LastUpdated    time.Time       `json:"-"`
}

type UsersData struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (u *User) IsOAuthUser() bool {
	return u.Provider != "local"
}

func (u *User) HasPassword() bool {
	return u.Password.Valid && u.Password.String != ""
}
