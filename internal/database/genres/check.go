package genres

func (db *GenreRepository) CheckGenreExistsByName(name string) (int, bool, error) {
	query := "SELECT id, EXISTS(SELECT 1 FROM genres WHERE name = $1) as exists FROM genres WHERE name = $1"
	var exists bool
	var id int
	err := db.Database.QueryRow(query, name).Scan(&id, &exists)
	if err != nil {
		return 0, false, err
	}
	return id, exists, nil
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
    SELECT EXISTS (
        SELECT 1 FROM movie_genres WHERE genre_id = $1 LIMIT 1
    )
    `

	err := db.Database.QueryRow(query, genreId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
