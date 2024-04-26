package posters

func (db *PosterRepo) CheckIfMovieHaveNoPosters(movieId int) (bool, error) {
	query := `SELECT COUNT(*) FROM posters WHERE movie_id = $1`

	var count int

	err := db.Database.QueryRow(query, movieId).Scan(&count)
	if err != nil {
		return false, err
	}

	return count != 0, nil
}
func (db *PosterRepo) CheckIfPosterExists(posterId int) (bool, error) {
	query := "SELECT COUNT(*) FROM posters WHERE id = $1"

	var count int
	err := db.Database.QueryRow(query, posterId).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
