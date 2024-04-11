package database

func (d *Database) UpdateProject(id int, imageUrl string, name string, category string, projectType string, year int, AgeCategory string, durationMinutes int, keywords string, desc string, director string, producer string) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := `UPDATE projects SET 
				image_url = $1,
				name = $2,
				category = $3,
				project_type = $4,
				year = $5,
				age_category = $6,
				duration_minutes = $7,
				keywords = $8,
				description = $9,
				director = $10,
				producer = $11,
				updated_at = CURRENT_TIMESTAMP
			  WHERE id = $12`

	_, err = tx.Exec(query, imageUrl, name, category, projectType, year, AgeCategory, durationMinutes, keywords, desc, director, producer, id)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
