package database

import "database/sql"

type Database struct {
	Database *sql.DB
}

type RegisterData struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
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
	IsAdmin     sql.NullInt16  `json:"is_admin"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	DeletedAt   sql.NullTime   `json:"deleted_at"`
}
type Movie struct {
	Id              int    `json:"id"`
	UserId          int    `json:"user_id"`
	ImageUrl        string `json:"image_url"`
	Name            string `json:"name"`
	Category        string `json:"category"`
	ProjectType     string `json:"project_type"`
	Year            int    `json:"year"`
	AgeCategory     string `json:"age_category"`
	DurationMinutes int    `json:"duration_minutes"`
	Keywords        string `json:"keywords"`
	Description     string `json:"description"`
	Director        string `json:"director"`
	Producer        string `json:"producer"`
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
