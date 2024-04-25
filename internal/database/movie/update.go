package movie

func (d *MovieRepository) UpdateMovieData(id int, name, description, director, producer string, year, durationMinutes int) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := `UPDATE movies SET name = $1, description = $2, director = $3, producer = $4, year = $5, duration_minutes = $6 WHERE id = $7`
	_, err = tx.Exec(q, name, description, director, producer, year, durationMinutes, id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (db *MovieRepository) MovieWasWatchedByUser(movieId int) (int, error) {
	tx, err := db.Database.Begin()
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	query := `UPDATE movies SET watch_count = watch_count + 1 WHERE id = $1 RETURNING watch_count`
	var watchCount int
	err = tx.QueryRow(query, movieId).Scan(&watchCount)
	if err != nil {
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}
	return watchCount, nil
}
func (db *MovieRepository) AddGenresToMovie(movieId int, genresId []int) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, genreID := range genresId {
		_, err = tx.Exec("INSERT INTO movie_genres (movie_id, genre_id) VALUES ($1, $2)", movieId, genreID)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (d *MovieRepository) UpdateMovieCategory(movieId, categoryId int) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := `UPDATE movies SET category_id = $1 WHERE id = $2`
	_, err = tx.Exec(q, categoryId, movieId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
func (d *MovieRepository) UpdateMovieAgeCategory(movieId, ageCategoryId int) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := `UPDATE movies SET age_category_id = $1 WHERE id = $2`
	_, err = tx.Exec(q, ageCategoryId, movieId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
func (db *MovieRepository) RemoveGenresFromMovie(movieId int, genresId []int) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, genreID := range genresId {
		_, err = tx.Exec("DELETE FROM movie_genres WHERE movie_id = $1 AND genre_id = $2", movieId, genreID)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
