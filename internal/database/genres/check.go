package genres

func (db *GenreRepository) CheckGenreExistsByName(name string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM genres WHERE name = $1)`
	var exists bool
	err := db.Database.QueryRow(query, name).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
func (db *GenreRepository) CheckGenreExistsById(id int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM genres WHERE id = $1)`
	var exists bool
	err := db.Database.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
func (db *GenreRepository) CheckGenreIsUsedInMovies(genreId int) (bool, error) {
	var exists bool

	query := `
        SELECT EXISTS(
            SELECT 1 FROM movies WHERE genre_id = $1
        )
    `

	err := db.Database.QueryRow(query, genreId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
