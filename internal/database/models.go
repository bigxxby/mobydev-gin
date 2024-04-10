package database

import (
	"database/sql"
)

type Database struct {
	Database *sql.DB
}

type RegisterData struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
	Role            string `json:"role" binding:"required"`
}

type LoginData struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
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
type Project struct {
	Id              int    `json:"id"`
	UserId          string `json:"userId"`
	ImageUrl        string `json:"imageUrl" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Category        string `json:"category" binding:"required"`
	ProjectType     string `json:"projectType" binding:"required"`
	Year            int    `json:"year" binding:"required"`
	AgeCategory     string `json:"ageCategory" binding:"required"`
	DurationMinutes int    `json:"durationMinutes" binding:"required"`
	Keywords        string `json:"keywords" binding:"required"`
	Description     string `json:"description" binding:"required"`
	Director        string `json:"director" binding:"required"`
	Producer        string `json:"producer"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

// type Movie struct {
// 	Id              int            `json:"id"`
// 	UserId          int            `json:"user_id"`          // для поддержки NULL
// 	ImageId         sql.NullInt64  `json:"image_id"`         // для поддержки NULL
// 	Name            sql.NullString `json:"name"`             // для поддержки NULL
// 	Category        sql.NullString `json:"category"`         // для поддержки NULL
// 	ProjectType     sql.NullString `json:"project_type"`     // для поддержки NULL
// 	Year            sql.NullInt64  `json:"year"`             // для поддержки NULL
// 	AgeCategory     sql.NullString `json:"age_category"`     // для поддержки NULL
// 	DurationMinutes sql.NullInt64  `json:"duration_minutes"` // для поддержки NULL
// 	Keywords        sql.NullString `json:"keywords"`         // для поддержки NULL
// 	Description     sql.NullString `json:"description"`      // для поддержки NULL
// 	Director        sql.NullString `json:"director"`         // для поддержки NULL
// 	Producer        sql.NullString `json:"producer"`         // для поддержки NULL
// }
