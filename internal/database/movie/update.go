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
