package movie

import (
	"database/sql"
	"project/internal/database/age"
	"project/internal/database/categories"
	"project/internal/database/genres"
	"project/internal/database/trend"
)

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
	Producer        string `json:"producer" binding:"required,"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`

	SeasonCount int `json:"seasonCount,omitempty" `
	SeriesCount int `json:"seriesCount,omitempty"`

	// WatchCount  int                    `json:"watchCount"`

	Categories []categories.Category `json:"category,omitempty"`

	Genres []genres.Genre `json:"genre,omitempty"`

	AgeCategories []age.AgeCategory `json:"ageCategory,omitempty"`

	Trend trend.Trend `json:"trend,omitempty"`

	IsFavorite bool `json:"isFavorite,omitempty"`
}

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
