package posters

func (db *PosterRepo) AddPostersMovieById(movieId int, posters [5]string) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	first := posters[0]
	sec := posters[1]
	third := posters[2]
	forth := posters[3]
	fifth := posters[4]

	q := `
		INSERT INTO posters (movie_id, main_poster, second_poster, third_poster, fourth_poster, fifth_poster) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = tx.Exec(q, movieId, first, sec, third, forth, fifth)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
