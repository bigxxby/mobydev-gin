package age

func (db *AgeRepository) DeleteAgeCategoryById(AgeId int) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := "DELETE from age_categories WHERE id = $1"

	_, err = tx.Exec(query, AgeId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
