package init

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"project/internal/database"
	"project/internal/database/age"
	"project/internal/database/episode"
	"project/internal/database/favorites"
	"project/internal/database/genres"
	"project/internal/database/movie"
	"project/internal/database/season"
	"project/internal/database/trend"

	"project/internal/database/categories"
	"project/internal/database/user"

	_ "github.com/lib/pq"
)

func CreateConnection() (*sql.DB, error) {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	log.Println("trying to connect...")
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database connected")
	return db, nil
}

func CreateDatabaseStruct() (*database.Database, error) {
	db, err := CreateConnection()
	if err != nil {
		return nil, err
	}
	database := database.Database{
		Database:             db,
		UserRepository:       &user.UserRepository{Database: db},
		MovieRepository:      &movie.MovieRepository{Database: db},
		FavoritesRepository:  &favorites.FavoritesRepository{Database: db},
		SeasonRepository:     &season.SeasonRepository{Database: db},
		EpisodeRepository:    &episode.EpisodeRepository{Database: db},
		TrendRepository:      &trend.TrendRepository{Database: db},
		CategoriesRepository: &categories.CategoryRepository{Database: db},
		GenreRepository:      &genres.GenreRepository{Database: db},
		AgeRepository:        &age.AgeRepository{Database: db},
	}
	return &database, err
}
