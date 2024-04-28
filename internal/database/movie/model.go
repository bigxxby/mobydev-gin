package movie

import (
	"database/sql"
)

type MovieRepository struct {
	Database *sql.DB
}

type Movie struct {
	Id     int    `json:"id"`
	UserId string `json:"userId"`
	Name   string `json:"name" binding:"required"`
	Year   int    `json:"year" binding:"required"`

	Poster [5]string `json:"posters"` //Maximum of 5 posters for one movie

	Category    string   `json:"category"`                  //one to many
	AgeCategory string   `json:"ageCategory"`               //one to many
	Genres      []string `json:"genres" binding:"required"` //many to many table

	CategoryId      int    `json:"categoryId" binding:"required"`
	AgeCategoryId   int    `json:"ageCategoryId" binding:"required"`
	DurationMinutes int    `json:"durationMinutes" binding:"required"`
	Keywords        string `json:"keywords" binding:"required"`
	Description     string `json:"description" binding:"required"`
	Director        string `json:"director" binding:"required"`
	Producer        string `json:"producer" binding:"required"`

	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	SeasonCount int    `json:"seasonCount" `
	SeriesCount int    `json:"seriesCount"`

	WatchCount int `json:"watchCount"`

	IsFavorite bool `json:"isFavorite"`
}
