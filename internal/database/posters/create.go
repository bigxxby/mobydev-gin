package posters

func (db *PosterRepo) AddPostersMovieById(movieId int, first, second, third, fourth, fifth string) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := `
		INSERT INTO posters (movie_id, main_poster, second_poster, third_poster, fourth_poster, fifth_poster) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = tx.Exec(q, movieId, first, second, third, fourth, fifth)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
