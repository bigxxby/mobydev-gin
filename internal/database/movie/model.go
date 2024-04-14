package movie

import "database/sql"

type MovieRepository struct {
	Database *sql.DB
}

type Movie struct {
	Id              int    `json:"id"`
	UserId          string `json:"userId"`
	ImageUrl        string `json:"imageUrl" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Year            int    `json:"year" binding:"required"`
	CategoryId      int    `json:"categoryId" binding:"required"`
	AgeCategoryId   int    `json:"ageCategoryId" binding:"required"`
	GenreId         int    `json:"genreId" binding:"required"`
	DurationMinutes int    `json:"durationMinutes" binding:"required"`
	Keywords        string `json:"keywords" binding:"required"`
	Description     string `json:"description" binding:"required"`
	Director        string `json:"director" binding:"required"`
	Producer        string `json:"producer" binding:"required"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}
