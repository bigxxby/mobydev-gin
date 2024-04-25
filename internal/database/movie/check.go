package movie

import (
	"database/sql"
)

func (db *MovieRepository) CheckMovieExistsById(movieId int) error {

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
