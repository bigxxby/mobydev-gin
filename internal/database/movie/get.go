package movie

func (db *MovieRepository) GetMovieById(id int) (*Movie, error) {
	stmt, err := db.Database.Prepare("SELECT * FROM movies WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var movie Movie

	err = stmt.QueryRow(id).Scan(&movie.Id, &movie.UserId, &movie.Name, &movie.Year,
		&movie.CategoryId, &movie.AgeCategoryId, &movie.WatchCount, &movie.DurationMinutes,
		&movie.Keywords, &movie.Description, &movie.Director, &movie.Producer,
		&movie.CreatedAt, &movie.UpdatedAt)
	if err != nil {
		return nil, err
	}

	err = db.Database.QueryRow("SELECT name FROM categories WHERE id = $1", movie.CategoryId).Scan(
		&movie.Category,
	)
	if err != nil {
		return nil, err
	}

	err = db.Database.QueryRow("SELECT name FROM age_categories WHERE id = $1", movie.AgeCategoryId).Scan(
		&movie.AgeCategory,
	)
	if err != nil {
		return nil, err
	}

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

	return &movie, nil
}

// get all movies
func (db *MovieRepository) GetMovies(userId int) ([]Movie, error) {
	stmt, err := db.Database.Prepare("SELECT * FROM movies")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie

	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.Id, &movie.UserId, &movie.Name, &movie.Year,
			&movie.CategoryId, &movie.AgeCategoryId, &movie.WatchCount, &movie.DurationMinutes,
			&movie.Keywords, &movie.Description, &movie.Director, &movie.Producer,
			&movie.CreatedAt, &movie.UpdatedAt)
		if err != nil {
			return nil, err
		}

		err = db.Database.QueryRow("SELECT name FROM categories WHERE id = $1", movie.CategoryId).Scan(
			&movie.Category,
		)
		if err != nil {
			return nil, err
		}

		err = db.Database.QueryRow("SELECT name FROM age_categories WHERE id = $1", movie.AgeCategoryId).Scan(
			&movie.AgeCategory,
		)
		if err != nil {
			return nil, err
		}
		if userId > 0 {

			var exists bool
			err = db.Database.QueryRow("SELECT EXISTS (SELECT 1 FROM favorites WHERE movie_id = $1 AND user_id=$2)", movie.Id, userId).Scan(
				&exists,
			)
			if err != nil {
				return nil, err
			}
			if exists {
				movie.IsFavorite = true
			}
		}

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
	stmt, err := db.Database.Prepare("SELECT * FROM movies LIMIT $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie

	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.Id, &movie.UserId, &movie.Name, &movie.Year,
			&movie.CategoryId, &movie.AgeCategoryId, &movie.WatchCount, &movie.DurationMinutes,
			&movie.Keywords, &movie.Description, &movie.Director, &movie.Producer,
			&movie.CreatedAt, &movie.UpdatedAt)
		if err != nil {
			return nil, err
		}

		err = db.Database.QueryRow("SELECT name FROM categories WHERE id = $1", movie.CategoryId).Scan(
			&movie.Category,
		)
		if err != nil {
			return nil, err
		}

		err = db.Database.QueryRow("SELECT name FROM age_categories WHERE id = $1", movie.AgeCategoryId).Scan(
			&movie.AgeCategory,
		)
		if err != nil {
			return nil, err
		}
		if userId > 0 {
			var exists bool
			err = db.Database.QueryRow("SELECT EXISTS (SELECT 1 FROM favorites WHERE movie_id = $1 AND user_id=$2)", movie.Id, userId).Scan(
				&exists,
			)
			if err != nil {
				return nil, err
			}
			if exists {
				movie.IsFavorite = true
			}
		}

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
