package movie

func (db *MovieRepository) SearchMovie(query string, userId int) ([]Movie, error) {
	searchQuery := "%" + query + "%" // Добавляем символы подстановки для частичного совпадения

	// Модифицированный SQL-запрос, который будет фильтровать фильмы по имени и ключевым словам
	sqlQuery := `
        SELECT m.id, m.user_id, m.name, m.year, m.category_id, m.age_category_id,
               m.watch_count, m.duration_minutes, m.keywords, m.description,
               m.director, m.producer, m.created_at, m.updated_at,
               c.name AS category_name, ac.name AS age_category_name,
               p.main_poster, p.second_poster, p.third_poster, p.fourth_poster, p.fifth_poster
        FROM movies m
        LEFT JOIN categories c ON m.category_id = c.id
        LEFT JOIN age_categories ac ON m.age_category_id = ac.id
        LEFT JOIN posters p ON m.id = p.movie_id
        WHERE LOWER(m.name) LIKE LOWER($1) OR LOWER(m.keywords) LIKE LOWER($1)
    `

	rows, err := db.Database.Query(sqlQuery, searchQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie

	for rows.Next() {
		var movie Movie
		rows.Scan(&movie.Id, &movie.UserId, &movie.Name, &movie.Year,
			&movie.CategoryId, &movie.AgeCategoryId, &movie.WatchCount, &movie.DurationMinutes,
			&movie.Keywords, &movie.Description, &movie.Director, &movie.Producer,
			&movie.CreatedAt, &movie.UpdatedAt, &movie.Category, &movie.AgeCategory,
			&movie.Poster[0], &movie.Poster[1], &movie.Poster[2],
			&movie.Poster[3], &movie.Poster[4])
		if err != nil {
			return nil, err
		}

		if userId > 0 {
			var exists bool
			err = db.Database.QueryRow("SELECT EXISTS (SELECT 1 FROM favorites WHERE movie_id = $1 AND user_id = $2)", movie.Id, userId).Scan(&exists)
			if err != nil {
				return nil, err
			}
			movie.IsFavorite = exists
		}

		// Оставшийся код для получения жанров и подсчета сезонов и серий остается без изменений

		// Добавление текущего фильма в список результатов
		movies = append(movies, movie)
	}

	return movies, nil
}

func (db *MovieRepository) GetSimilarMoviesLimit5(movieKeywords string, excludeMovieID int) ([]Movie, error) {
	var similarMovies []Movie

	query := `
        SELECT m.id, m.name, p.main_poster
        FROM movies m
        JOIN posters p ON m.id = p.movie_id
        WHERE m.keywords LIKE '%' || $1 || '%'
        AND m.id != $2
        LIMIT 5
    `

	rows, err := db.Database.Query(query, movieKeywords, excludeMovieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.Id, &movie.Name, &movie.Poster[0]); err != nil {
			return nil, err
		}
		similarMovies = append(similarMovies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return similarMovies, nil
}
