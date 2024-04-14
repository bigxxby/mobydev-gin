package datasets

import (
	"database/sql"
	"log"
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
	err = insertTestTrends(db)
	if err != nil {
		return err
	}
	err = insertTestFavorites(db)
	if err != nil {
		return err
	}

	return nil
}
func DropTables(db *sql.DB) error {
	tables := []string{"episodes", "seasons", "movies", "users", "trends", "favorites", "categories", "age_categories", "genres"}

	for _, table := range tables {
		_, err := db.Exec("DROP TABLE IF EXISTS " + table + " CASCADE")
		if err != nil {
			log.Fatalf("Failed to drop table %s: %v", table, err)
			return err
		}
	}

	log.Println("All tables dropped successfully")
	return nil
}
