package movie

func (d *MovieRepository) UpdateMovie(id int, imageUrl string, name string, year int, categoryId int, ageCategoryId int, genreId int, durationMinutes int, keywords string, desc string, director string, producer string) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := `UPDATE movies SET 
				image_url = $1,
				name = $2,
				year = $3,
				category_id = $4,
				age_category_id = $5,
				genre_id = $6,
				duration_minutes = $7,
				keywords = $8,
				description = $9,
				director = $10,
				producer = $11,
				updated_at = CURRENT_TIMESTAMP
			  WHERE id = $12`

	_, err = tx.Exec(query, imageUrl, name, year, categoryId, ageCategoryId, genreId, durationMinutes, keywords, desc, director, producer, id)
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
	return watchCount, nil // Возврат текущего количества просмотров фильма
}
