package init

import (
	"database/sql"
	"log"
	"project/internal/database"
	"project/internal/database/episode"
	"project/internal/database/favorites"
	"project/internal/database/movie"
	"project/internal/database/season"
	"project/internal/database/trend"
	"project/internal/database/user"

	_ "github.com/lib/pq"
)

func CreateConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=db port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")

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
		Database:            db,
		UserRepository:      &user.UserRepository{Database: db},
		MovieRepository:     &movie.MovieRepository{Database: db},
		FavoritesRepository: &favorites.FavoritesRepository{Database: db},
		SeasonRepository:    &season.SeasonRepository{Database: db},
		EpisodeRepository:   &episode.EpisodeRepository{Database: db},
		TrendRepository:     &trend.TrendRepository{Database: db},
	}
	return &database, err
}
