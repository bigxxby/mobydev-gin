package genres

import (
	"database/sql"
)

func (db *GenreRepository) DeleteGenreById(id int) error {
	query := "DELETE FROM genres WHERE id = $1"
	_, err := db.Database.Exec(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return sql.ErrNoRows
		}
		return err
	}
	return nil
}
