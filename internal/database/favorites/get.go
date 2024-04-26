package favorites

func (db *FavoritesRepository) GetFavoritesByUserId(id int) ([]Favorite, error) {
	var favorites []Favorite

	query := `
		SELECT f.id, f.user_id, f.movie_id, f.added_at, m.name, m.year, m.category_id, m.age_category_id, m.watch_count, 
			   m.duration_minutes, m.keywords, m.description, m.director, m.producer, m.created_at, m.updated_at
		FROM favorites f
		JOIN movies m ON f.movie_id = m.id
		WHERE f.user_id = $1
	`

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
			&favorite.Movie.Name,
			&favorite.Movie.Year,
			&favorite.Movie.CategoryId,
			&favorite.Movie.AgeCategoryId,
			&favorite.Movie.WatchCount,
			&favorite.Movie.DurationMinutes,
			&favorite.Movie.Keywords,
			&favorite.Movie.Description,
			&favorite.Movie.Director,
			&favorite.Movie.Producer,
			&favorite.Movie.CreatedAt,
			&favorite.Movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
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
