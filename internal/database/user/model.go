package user

import "database/sql"

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
