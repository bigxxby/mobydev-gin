package routes

import (
	"project/internal/database/movie"
	"time"
)

type SignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type SignUpRequest struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

type SignInResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type DefaultMessageResponse struct {
	Message string `json:"message"`
}
type ChangePasswordRequest struct {
	OldPassword     string `json:"oldPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}
type VerifyCodeRequest struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required"`
}
type SendCodeRequest struct {
	Email string `json:"email" binding:"required"`
}
type ManyMoviesResponse struct {
	Movies []movie.Movie `json:"movies"`
}
type MovieResponse struct {
	Movies movie.Movie `json:"movie"`
}
type MovieCreateRequest struct {
	Name            string   `json:"name" binding:"required"`
	Year            int      `json:"year" binding:"required"`
	Genres          []string `json:"genres" binding:"required"` //many to many table
	CategoryId      int      `json:"categoryId" binding:"required"`
	AgeCategoryId   int      `json:"ageCategoryId" binding:"required"`
	DurationMinutes int      `json:"durationMinutes" binding:"required"`
	Keywords        string   `json:"keywords" binding:"required"`
	Description     string   `json:"description" binding:"required"`
	Director        string   `json:"director" binding:"required"`
	Producer        string   `json:"producer" binding:"required"`
}
type MovieDataRequest struct {
	Name            string `json:"name" binding:"required"`
	Description     string `json:"description" binding:"required"`
	Year            int    `json:"year" binding:"required"`
	DurationMinutes int    `json:"durationMinutes" binding:"required"`
	Director        string `json:"director" binding:"required"`
	Producer        string `json:"producer" binding:"required"`
	Keywords        string `json:"keywords" binding:"required"`
}
type MovieGenresRequest struct {
	Genres []string `json:"genres" binding:"required"`
}
type AgeCategoryRequest struct {
	Name   string `json:"name" binding:"required"`
	Note   string `json:"note"`
	MinAge int    `json:"min_age" binding:"required"`
	MaxAge int    `json:"max_age" binding:"required"`
}
type CategoryRequest struct {
	Name        string `json:"category_name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
type EpisodeRequest struct {
	URL             string `json:"url" binding:"required"`
	EpisodeNumber   int    `json:"episode_number" binding:"required"`
	Name            string `json:"name" binding:"required"`
	DurationMinutes int    `json:"duration_minutes" binding:"required"`
	Description     string `json:"description"`
	ReleaseDate     string `json:"release_date" binding:"required"`
}
type EpisodeRequestBody struct {
	Episodes []EpisodeRequest `json:"episodes"`
}
type UserProfileResponse struct {
	ID    int       `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Phone string    `json:"phone"`
	Role  string    `json:"role"`
	Dot   time.Time `json:"dot"`
}
type UserProfileRequest struct {
	Name  string    `json:"name"`
	Phone string    `json:"phone"`
	Dot   time.Time `json:"dot"`
}
type SeasonBodyRequest struct {
	SeasonNumber int    `json:"season_number" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	ReleaseDate  string `json:"release_date" binding:"required"`
}
type SeasonsBodyRequest struct {
	Seasons []SeasonBodyRequest `json:"seasons" binding:"required"`
}
