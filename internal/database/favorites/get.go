package favorites

func (db *FavoritesRepository) GetFavoritesByUserId(id int) ([]Favorite, error) {
	var favorites []Favorite

	query := `SELECT * FROM favorites WHERE user_id = $1`

	rows, err := db.Database.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var favorite Favorite

		err := rows.Scan(
			&favorite.ID,
			&favorite.UserID,
			&favorite.MovieID,
			&favorite.AddedAt,
		)

		if err != nil {
			return nil, err
		}
		favorites = append(favorites, favorite)

	}

	return favorites, nil
}
