package database

import (
	"database/sql"
	"project/internal/database/categories"
	"project/internal/database/episode"
	"project/internal/database/favorites"
	"project/internal/database/movie"
	"project/internal/database/season"
	"project/internal/database/trend"
	"project/internal/database/user"
)

type Database struct {
	Database             *sql.DB
	UserRepository       *user.UserRepository
	MovieRepository      *movie.MovieRepository
	SeasonRepository     *season.SeasonRepository
	EpisodeRepository    *episode.EpisodeRepository
	TrendRepository      *trend.TrendRepository
	FavoritesRepository  *favorites.FavoritesRepository
	CategoriesRepository *categories.CategoryRepository
}
