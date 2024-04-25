package categories

func (db *CategoryRepository) GetCategoryById(id int) (*Category, error) {
	var category Category

	query := `SELECT * FROM categories WHERE id = $1`

	err := db.Database.QueryRow(query, id).Scan(
		&category.ID,
		&category.UserID,
		&category.Name,
		&category.Description,
	)

	if err != nil {
		return nil, err
	}

	return &category, nil
}
func (db *CategoryRepository) GetCategoryByName(name string) (*Category, error) {
	var category Category

	query := `SELECT * FROM categories WHERE name = $1`

	err := db.Database.QueryRow(query, name).Scan(
		&category.ID,
		&category.UserID,
		&category.Name,
		&category.Description,
	)

	if err != nil {
		return nil, err
	}

	return &category, nil
}
func (db *CategoryRepository) GetCategories() ([]Category, error) {
	stmt, err := db.Database.Prepare("SELECT * FROM categories ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category

	for rows.Next() {
		var category Category
		err := rows.Scan(
			&category.ID,
			&category.UserID,
			&category.Name,
			&category.Description,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
