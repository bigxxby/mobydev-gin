package movie

import "database/sql"

type MovieRepository struct {
	Database *sql.DB
	GET      GET
}
type GET interface {
	GetMovies(limit int) ([]Movie, error)
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
	Producer        string `json:"producer" binding:"required"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}
