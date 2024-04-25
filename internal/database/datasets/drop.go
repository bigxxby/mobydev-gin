package datasets

import (
	"database/sql"
	"log"
)

func DropTables(db *sql.DB) error {
	tables := []string{"episodes", "seasons", "movies", "users", "favorites", "categories", "age_categories", "codes", "movie_genres", "genres"}

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
