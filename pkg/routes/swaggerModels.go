package routes

import "project/internal/database/movie"

type SignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type SignUpRequest struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

// SignInResponse представляет структуру ответа на успешную аутентификацию
type SignInResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// routes.DefaultMessageResponse
type DefaultMessageResponse struct {
	Message string `json:"message"`
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
