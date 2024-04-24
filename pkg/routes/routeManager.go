package routes

import (
	"project/internal/database"
	i "project/internal/init"
	"project/pkg/routes/api/age"
	"project/pkg/routes/api/auth"
	"project/pkg/routes/api/categories"
	"project/pkg/routes/api/episodes"
	"project/pkg/routes/api/favorites"
	"project/pkg/routes/api/genres"
	"project/pkg/routes/api/movies"
	"project/pkg/routes/api/seasons"
	"project/pkg/routes/api/users"
)

type Manager struct {
	DB              *database.Database
	MoviesRoute     movies.MoviesRoute
	UsersRoute      users.UsersRoute
	EpisodesRoute   episodes.EpisodesRoute
	SeasonsRoute    seasons.SeasonsRoute
	FavoritesRoute  favorites.FavoritesRoute
	AuthRoute       auth.AuthRoute
	CategoriesRoute categories.CategoriesRoute
	GenreRoute      genres.GenreRoute
	AgeRoute        age.AgeRoute
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
	authRoute := auth.AuthRoute{DB: db}
	categoriesRoute := categories.CategoriesRoute{DB: db}
	genreRoute := genres.GenreRoute{DB: db}
	ageRoute := age.AgeRoute{DB: db}

	manager.DB = db
	manager.MoviesRoute = moviesRoute
	manager.UsersRoute = usersRoute
	manager.SeasonsRoute = seasonsRoute
	manager.FavoritesRoute = favoritesRoute
	manager.EpisodesRoute = episodesRoute
	manager.AuthRoute = authRoute
	manager.CategoriesRoute = categoriesRoute
	manager.GenreRoute = genreRoute
	manager.AgeRoute = ageRoute
	if err != nil {
		return nil, err
	}
	return &manager, nil
}
