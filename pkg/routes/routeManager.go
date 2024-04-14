package routes

import (
	"project/internal/database"
	i "project/internal/init"
	"project/pkg/routes/api/auth"
	"project/pkg/routes/api/categories"
	"project/pkg/routes/api/episodes"
	"project/pkg/routes/api/favorites"
	"project/pkg/routes/api/genres"
	"project/pkg/routes/api/movies"
	"project/pkg/routes/api/seasons"
	"project/pkg/routes/api/trends"
	"project/pkg/routes/api/users"
)

type Manager struct {
	DB              *database.Database
	MoviesRoute     movies.MoviesRoute
	UsersRoute      users.UsersRoute
	EpisodesRoute   episodes.EpisodesRoute
	SeasonsRoute    seasons.SeasonsRoute
	FavoritesRoute  favorites.FavoritesRoute
	TrendsRoute     trends.TrendsRoute
	AuthRoute       auth.AuthRoute
	CategoriesRoute categories.CategoriesRoute
	GenreRoute      genres.GenreRoute
}

func Init() (*Manager, error) {
	manager := Manager{}
	var err error
	db, err := i.CreateDatabaseStruct()

	moviesRoute := movies.MoviesRoute{DB: db}
	usersRoute := users.UsersRoute{DB: db}
	episodesRoute := episodes.EpisodesRoute{DB: db}
	seasonsRoute := seasons.SeasonsRoute{DB: db}
	favoritesRoute := favorites.FavoritesRoute{DB: db}
	trendsRoute := trends.TrendsRoute{DB: db}
	authRoute := auth.AuthRoute{DB: db}
	categoriesRoute := categories.CategoriesRoute{DB: db}
	genreRoute := genres.GenreRoute{DB: db}

	manager.DB = db
	manager.MoviesRoute = moviesRoute
	manager.UsersRoute = usersRoute
	manager.SeasonsRoute = seasonsRoute
	manager.FavoritesRoute = favoritesRoute
	manager.TrendsRoute = trendsRoute
	manager.EpisodesRoute = episodesRoute
	manager.AuthRoute = authRoute
	manager.CategoriesRoute = categoriesRoute
	manager.GenreRoute = genreRoute

	if err != nil {
		return nil, err
	}
	return &manager, nil
}
