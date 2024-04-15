package age

func (db *AgeRepository) CreateAgeCategory(userId int, name, note string, minAge, maxAge int) (int, error) {
	var noteStr string

	query := `INSERT INTO age_categories (user_id, name, note, min_age, max_age) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int
	err := db.Database.QueryRow(query, userId, name, noteStr, minAge, maxAge).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
