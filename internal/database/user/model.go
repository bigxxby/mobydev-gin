package user

import (
	"database/sql"
	"time"
)

type UserRepository struct {
	Database *sql.DB
}
type User struct {
	Id          int            `json:"id"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	Name        sql.NullString `json:"name"`
	Phone       sql.NullString `json:"phone"`
	DateOfBirth sql.NullTime   `json:"date_of_birth"`
	Role        string         `json:"role"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	DeletedAt   sql.NullTime   `json:"deleted_at"`
}
type UserShort struct {
	Name        string `json:"name" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
}
type UserJson struct {
	Id          int       `json:"id"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Name        string    `json:"name,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	DateOfBirth string    `json:"date_of_birth,omitempty"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
