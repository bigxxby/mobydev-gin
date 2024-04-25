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

	//options
	Category    string   `json:"category"`
	AgeCategory string   `json:"ageCategory"`
	Genres      []string `json:"genres" binding:"required"` //many to many table

	CategoryId    int `json:"categoryId" binding:"required"`
	AgeCategoryId int `json:"ageCategoryId" binding:"required"`

	DurationMinutes int    `json:"durationMinutes" binding:"required"`
	Keywords        string `json:"keywords" binding:"required"`
	Description     string `json:"description" binding:"required"`
	Director        string `json:"director" binding:"required"`
	Producer        string `json:"producer" binding:"required"`

	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	SeasonCount int    `json:"seasonCount,omitempty" `
	SeriesCount int    `json:"seriesCount,omitempty"`

	WatchCount int `json:"watchCount"`

	IsFavorite bool `json:"isFavorite"`
}

// movie that is created (no genres required)

// type MovieMain struct {
// 	Id              int    `json:"id"`
// 	UserId          string `json:"userId"`
// 	ImageUrl        string `json:"imageUrl" binding:"required"`
// 	Name            string `json:"name" binding:"required"`
// 	Year            int    `json:"year" binding:"required"`
// 	CategoryId      int    `json:"categoryId" binding:"required"`
// 	AgeCategoryId   int    `json:"ageCategoryId" binding:"required"`
// 	GenreId         int    `json:"genreId" binding:"required"`
// 	DurationMinutes int    `json:"durationMinutes" binding:"required"`
// 	Keywords        string `json:"keywords" binding:"required"`
// 	Description     string `json:"description" binding:"required"`
// 	Director        string `json:"director" binding:"required"`
// 	Producer        string `json:"producer" binding:"required"`
// 	CreatedAt       string `json:"createdAt"`
// 	UpdatedAt       string `json:"updatedAt"`

// 	SeasonCount int `json:"seasonCount"`
// 	SeriesCount int `json:"seriesCount"`

// 	// WatchCount  int                    `json:"watchCount"`

// 	Categories []categories.Category `json:"category"`

// 	Genres []genres.Genre `json:"genre"`

// 	AgeCategories []age.AgeCategory `json:"ageCategory"`

// 	Trend trend.Trend `json:"trend"`

// 	IsFavorite bool `json:"isFavorite"`
// }
