package categories

// updates cat. by id
func (db *CategoryRepository) UpdateCategory(categoryId int, newName, newDescription string) (int, error) {

	tx, err := db.Database.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	query := `
        UPDATE categories
        SET name = $1, description = $2
        WHERE id = $3
    `
	_, err = tx.Query(query, newName, newDescription, categoryId)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return categoryId, nil
}
