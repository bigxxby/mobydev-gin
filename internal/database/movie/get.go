package movie

func (db *MovieRepository) GetMovieById(userId, movieId int) (*Movie, error) {
	// Запрос для выборки всех данных о фильме с заданным ID
	query := `
        SELECT m.id, m.user_id, m.name, m.year, m.category_id, m.age_category_id,
               m.watch_count, m.duration_minutes, m.keywords, m.description,
               m.director, m.producer, m.created_at, m.updated_at,
               c.name AS category_name, ac.name AS age_category_name,
               p.main_poster, p.second_poster, p.third_poster, p.fourth_poster, p.fifth_poster
        FROM movies m
        LEFT JOIN categories c ON m.category_id = c.id
        LEFT JOIN age_categories ac ON m.age_category_id = ac.id
        LEFT JOIN posters p ON m.id = p.movie_id
        WHERE m.id = $1
    `

	row := db.Database.QueryRow(query, movieId)

	var movie Movie

	row.Scan(&movie.Id, &movie.UserId, &movie.Name, &movie.Year,
		&movie.CategoryId, &movie.AgeCategoryId, &movie.WatchCount, &movie.DurationMinutes,
		&movie.Keywords, &movie.Description, &movie.Director, &movie.Producer,
		&movie.CreatedAt, &movie.UpdatedAt, &movie.Category, &movie.AgeCategory,
		&movie.Poster[0], &movie.Poster[1], &movie.Poster[2],
		&movie.Poster[3], &movie.Poster[4])

	genreQuery := `
        SELECT g.name
        FROM genres g
        INNER JOIN movie_genres mg ON g.id = mg.genre_id
        WHERE mg.movie_id = $1
    `
	genreRows, err := db.Database.Query(genreQuery, movie.Id)
	if err != nil {
		return nil, err
	}
	defer genreRows.Close()

	genreRows, err = db.Database.Query("SELECT g.name FROM genres g INNER JOIN movie_genres mg ON g.id = mg.genre_id WHERE mg.movie_id = $1", movie.Id)
	if err != nil {
		return nil, err
	}
	defer genreRows.Close()

	var genresArr []string
	for genreRows.Next() {
		var genre string
		err := genreRows.Scan(&genre)
		if err != nil {
			return nil, err
		}
		genresArr = append(genresArr, genre)
	}
	movie.Genres = genresArr

	var isFavorite int
	err = db.Database.QueryRow("SELECT COUNT(*) FROM favorites WHERE user_id = $1 AND movie_id = $2", userId, movieId).Scan(&isFavorite)
	if err != nil {
		return nil, err
	}
	movie.IsFavorite = isFavorite > 0

	return &movie, nil
}

func (db *MovieRepository) GetMovies(userId int) ([]Movie, error) {
	query := `
        SELECT m.id, m.user_id, m.name, m.year, m.category_id, m.age_category_id,
               m.watch_count, m.duration_minutes, m.keywords, m.description,
               m.director, m.producer, m.created_at, m.updated_at,
               c.name AS category_name, ac.name AS age_category_name,
               p.main_poster, p.second_poster, p.third_poster, p.fourth_poster, p.fifth_poster
        FROM movies m
        LEFT JOIN categories c ON m.category_id = c.id
        LEFT JOIN age_categories ac ON m.age_category_id = ac.id
        LEFT JOIN posters p ON m.id = p.movie_id
    `
	rows, err := db.Database.Query(query)
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

		// Получение жанров фильма
		genreRows, err := db.Database.Query("SELECT g.name FROM genres g INNER JOIN movie_genres mg ON g.id = mg.genre_id WHERE mg.movie_id = $1", movie.Id)
		if err != nil {
			return nil, err
		}
		defer genreRows.Close()

		var genresArr []string
		for genreRows.Next() {
			var genre string
			err := genreRows.Scan(&genre)
			if err != nil {
				return nil, err
			}
			genresArr = append(genresArr, genre)
		}
		movie.Genres = genresArr

		movies = append(movies, movie)
	}

	return movies, nil
}

// returns movies include limit
func (db *MovieRepository) GetMoviesLimit(limit int, userId int) ([]Movie, error) {
	query := `
        SELECT m.id, m.user_id, m.name, m.year, m.category_id, m.age_category_id,
               m.watch_count, m.duration_minutes, m.keywords, m.description,
               m.director, m.producer, m.created_at, m.updated_at,
               c.name AS category_name, ac.name AS age_category_name,
               p.main_poster, p.second_poster, p.third_poster, p.fourth_poster, p.fifth_poster
        FROM movies m
        LEFT JOIN categories c ON m.category_id = c.id
        LEFT JOIN age_categories ac ON m.age_category_id = ac.id
        LEFT JOIN posters p ON m.id = p.movie_id
        LIMIT $1
    `
	rows, err := db.Database.Query(query, limit)
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

		// Получение жанров фильма
		genreRows, err := db.Database.Query("SELECT g.name FROM genres g INNER JOIN movie_genres mg ON g.id = mg.genre_id WHERE mg.movie_id = $1", movie.Id)
		if err != nil {
			return nil, err
		}
		defer genreRows.Close()

		var genresArr []string
		for genreRows.Next() {
			var genre string
			err := genreRows.Scan(&genre)
			if err != nil {
				return nil, err
			}
			genresArr = append(genresArr, genre)
		}
		movie.Genres = genresArr

		movies = append(movies, movie)
	}

	return movies, nil
}

func (db *MovieRepository) GetGenresByMovieId(movieId int) ([]int, error) {
	stmt, err := db.Database.Prepare("SELECT genre_id FROM movie_genres WHERE movie_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(movieId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genresId []int

	for rows.Next() {
		var genreID int
		err := rows.Scan(&genreID)
		if err != nil {
			return nil, err
		}
		genresId = append(genresId, genreID)
	}

	return genresId, nil
}
