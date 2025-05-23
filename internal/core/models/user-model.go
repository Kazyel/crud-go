package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
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
	Name        pgtype.Text      `json:"name,omitempty"`
	Email       pgtype.Text      `json:"email,omitempty"`
	Password    pgtype.Text      `json:"-"`
	LastUpdated pgtype.Timestamp `json:"-"`
}
