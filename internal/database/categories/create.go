package categories

func (db *CategoryRepository) CreateCategory(userID int, name, description string) (int, error) {

	tx, err := db.Database.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	query := `
        INSERT INTO categories (user_id, name ,  description)
        VALUES ($1, $2, $3) RETURNING id
    `

	var categoryId int
	err = tx.QueryRow(query, userID, name, description).Scan(&categoryId)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return categoryId, nil
}
