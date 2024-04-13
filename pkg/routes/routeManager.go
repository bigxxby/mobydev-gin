package routes

import (
	"project/internal/database"
	i "project/internal/init"
	"project/pkg/routes/api/auth"
	"project/pkg/routes/api/episodes"
	"project/pkg/routes/api/favorites"
	"project/pkg/routes/api/movies"
	"project/pkg/routes/api/seasons"
	"project/pkg/routes/api/trends"
	"project/pkg/routes/api/users"
)

type Manager struct {
	DB             *database.Database
	MoviesRoute    movies.MoviesRoute
	UsersRoute     users.UsersRoute
	EpisodesRoute  episodes.EpisodesRoute
	SeasonsRoute   seasons.SeasonsRoute
	FavoritesRoute favorites.FavoritesRoute
	TrendsRoute    trends.TrendsRoute
	AuthRoute      auth.AuthRoute
}

func Init() (*Manager, error) {
	manager := Manager{}
	var err error
	db, err := i.CreateDatabaseStruct()

	moviesRoute := movies.MoviesRoute{db}
	usersRoute := users.UsersRoute{db}
	episodesRoute := episodes.EpisodesRoute{db}
	seasonsRoute := seasons.SeasonsRoute{db}
	favoritesRoute := favorites.FavoritesRoute{db}
	trendsRoute := trends.TrendsRoute{db}
	authRoute := auth.AuthRoute{db}

	manager.DB = db
	manager.MoviesRoute = moviesRoute
	manager.UsersRoute = usersRoute
	manager.SeasonsRoute = seasonsRoute
	manager.FavoritesRoute = favoritesRoute
	manager.TrendsRoute = trendsRoute
	manager.EpisodesRoute = episodesRoute
	manager.AuthRoute = authRoute

	if err != nil {
		return nil, err
	}
	return &manager, nil
}
