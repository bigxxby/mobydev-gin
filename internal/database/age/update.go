package age

func (db *AgeRepository) UpdateAgeById(id int, name string, note string, minAge int, maxAge int) error {

	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		UPDATE age_categories
		SET name = $1, note = $2, min_age = $3, max_age = $4
		WHERE id = $5
	`

	_, err = tx.Exec(query, name, note, minAge, maxAge, id)
	if err != nil {

		return err
	}

	if err := tx.Commit(); err != nil {

		return err
	}

	return nil
}
