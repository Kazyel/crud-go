package models

import (
	"time"
)

type User struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"-" db:"password_hash"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	LastUpdated time.Time `json:"last_updated" db:"last_updated"`
}

type UserUpdate struct {
	Name        *string   `json:"name,omitempty"`
	Email       *string   `json:"email,omitempty"`
	Password    *string   `json:"-"`
	LastUpdated time.Time `json:"-"`
}

type UsersData struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
