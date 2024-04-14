package genres

func (db *GenreRepository) CreateGenre(userId int, name string, description string) (int, error) {
	query := `INSERT INTO genres (user_id, name, description) VALUES ($1, $2, $3)`

	result, err := db.Database.Query(query, userId, name, description)
	if err != nil {
		return 0, err
	}
	var genreId int
	result.Scan(&genreId)

	return genreId, nil
}
