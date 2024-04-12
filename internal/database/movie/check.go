package movie

import (
	"database/sql"
	"log"
)

func (db *MovieRepository) CheckMovieExistsById(movieId int) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	stmt, err := db.Database.Prepare("SELECT * FROM movies WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(movieId)
	if err != nil {
		return err
	}
	defer rows.Close()

	exists := rows.Next()
	if !exists {
		return sql.ErrNoRows
	}

	return nil
}
