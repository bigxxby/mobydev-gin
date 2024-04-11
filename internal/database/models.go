package database

import (
	"database/sql"
	"time"
)

type Database struct {
	Database *sql.DB
}

type RegisterData struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
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
type Movie struct {
	Id              int    `json:"id"`
	UserId          string `json:"userId"`
	ImageUrl        string `json:"imageUrl" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Category        string `json:"category" binding:"required"`
	MovieType       string `json:"projectType" binding:"required"`
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

type Season struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	MovieID      int       `json:"movie_id"`
	SeasonNumber int       `json:"season_number"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	ReleaseDate  time.Time `json:"release_date"`
}
type Episode struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	URL             string    `json:"url"`
	SeasonID        int       `json:"season_id"`
	EpisodeNumber   int       `json:"episode_number"`
	Name            string    `json:"name"`
	DurationMinutes int       `json:"duration_minutes"`
	ReleaseDate     time.Time `json:"release_date"`
	Description     string    `json:"description"`
}
type Trend struct {
	ID         int       `json:"id"`
	MovieID    int       `json:"movie_id"`
	TrendDate  time.Time `json:"trend_date"`
	TrendValue int       `json:"trend_value"`
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
