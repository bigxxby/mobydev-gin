package age

func (db *AgeRepository) CheckAgeCategoryExistsByName(name string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM age_categories WHERE name = $1)"
	var exists bool
	err := db.Database.QueryRow(query, name).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
