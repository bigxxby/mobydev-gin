package favorites

func (db *FavoritesRepository) GetFavoritesByUserId(id int) ([]Favorite, error) {
	var favorites []Favorite

	query := `
		SELECT f.id, f.user_id, f.movie_id, f.added_at, m.name, p.main_poster
		FROM favorites f
		LEFT JOIN movies m ON f.movie_id = m.id
		LEFT JOIN posters p ON f.movie_id = p.movie_id
		WHERE f.user_id = $1
	`

	rows, err := db.Database.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var favorite Favorite

		rows.Scan(
			&favorite.ID,
			&favorite.UserID,
			&favorite.MovieID,
			&favorite.AddedAt,
			&favorite.Movie.Name,
			&favorite.Movie.MainPoster,
		)
		favorite.Movie.Id = favorite.MovieID

		favorites = append(favorites, favorite)
	}

	return favorites, nil
}

// func (db *FavoritesRepository) GetFavoriteMoviesByUserId(id int) ([]movie.Movie, error) {
// 	var favorites []movie.Movie

// 	// SQL запрос
// 	query := `
//         SELECT m.*
//         FROM movies m
//         JOIN favorites f ON m.id = f.movie_id
//         WHERE f.user_id = $1
//     `

// 	// Выполнение запроса
// 	rows, err := db.Database.Query(query, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	// Обработка результатов запроса
// 	for rows.Next() {
// 		var favorite movie.Movie
// 		// Сканирование данных в структуру Movie
// 		err := rows.Scan(&favorite.Id, &favorite.UserId, &favorite.ImageUrl, &favorite.Name, &favorite.Year /* остальные поля movie.Movie */)
// 		if err != nil {
// 			return nil, err
// 		}
// 		favorites = append(favorites, favorite)
// 	}

// 	// Обработка возможных ошибок после завершения итерации по результатам запроса
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return favorites, nil
// }
