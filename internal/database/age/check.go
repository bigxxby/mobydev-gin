package age

func (db *AgeRepository) CheckAgeCategoryExistsByName(name string) (bool, error) {
	tx, err := db.Database.Begin()
	if err != nil {

		return false, err
	}
	defer tx.Rollback()
	query := "SELECT EXISTS(SELECT 1 FROM age_categories WHERE name = $1)"
	var exists bool
	err = tx.QueryRow(query, name).Scan(&exists)
	if err != nil {
		return false, err
	}
	err = tx.Commit()
	if err != nil {
		return false, err
	}
	return exists, nil
}
func (db *AgeRepository) CheckAgeCategoryExistsId(ageId int) (bool, error) {
	tx, err := db.Database.Begin()
	if err != nil {

		return false, err
	}
	defer tx.Rollback()
	query := "SELECT EXISTS(SELECT 1 FROM age_categories WHERE id = $1)"
	var exists bool
	err = tx.QueryRow(query, ageId).Scan(&exists)
	if err != nil {
		return false, err
	}
	err = tx.Commit()
	if err != nil {
		return false, err
	}
	return exists, nil
}
