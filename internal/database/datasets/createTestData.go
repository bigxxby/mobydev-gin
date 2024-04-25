package datasets

import (
	"project/internal/database"
)

func CreateTestData(db *database.Database) error {
	err := insertTestUsers(db)
	if err != nil {
		return err
	}
	err = insertTestGenres(db)
	if err != nil {
		return err
	}
	err = insertTestAgeCategoires(db)
	if err != nil {
		return err
	}
	err = insertTestCategories(db)
	if err != nil {
		return err
	}

	err = insertTestMovies(db)
	if err != nil {
		return err
	}

	err = insertTestSeasons(db)
	if err != nil {
		return err
	}

	err = insertTestEpisodes(db)
	if err != nil {
		return err
	}

	err = insertTestFavorites(db)
	if err != nil {
		return err
	}
	err = insertTestMovieGenres(db)
	if err != nil {
		return err
	}

	return nil
}
