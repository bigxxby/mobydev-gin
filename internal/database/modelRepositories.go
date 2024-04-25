package database

import (
	"database/sql"
	"project/internal/database/age"
	"project/internal/database/categories"
	"project/internal/database/episode"
	"project/internal/database/favorites"
	"project/internal/database/genres"
	"project/internal/database/movie"
	"project/internal/database/posters"
	"project/internal/database/season"
	"project/internal/database/user"
)

type Database struct {
	Database             *sql.DB
	UserRepository       *user.UserRepository
	MovieRepository      *movie.MovieRepository
	SeasonRepository     *season.SeasonRepository
	EpisodeRepository    *episode.EpisodeRepository
	FavoritesRepository  *favorites.FavoritesRepository
	CategoriesRepository *categories.CategoryRepository
	GenreRepository      *genres.GenreRepository
	AgeRepository        *age.AgeRepository
	PosterRepo           *posters.PosterRepo
}
