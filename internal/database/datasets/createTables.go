package datasets

import (
	"project/internal/database"
)

func CreateTables(db *database.Database) error {
	err := CreateUsersTable(db)
	if err != nil {
		return err
	}
	err = CreateGenresTable(db)
	if err != nil {
		return err
	}
	err = CreateCategoriesTable(db)
	if err != nil {
		return err
	}
	err = CreateAgeCategoriesTable(db)
	if err != nil {
		return err
	}
	err = CreateMoviesTable(db)
	if err != nil {
		return err
	}
	err = CreateSeasonsTable(db)
	if err != nil {
		return err
	}
	err = CreateEpisodesTable(db)
	if err != nil {
		return err
	}

	err = CreateFavoritesTable(db)
	if err != nil {
		return err
	}
	err = CreateCodesTable(db)
	if err != nil {
		return err
	}
	err = CreateMovieGenresTable(db)
	if err != nil {
		return err
	}
	err = CreateTablePosters(db)
	if err != nil {
		return err
	}
	return nil
}
