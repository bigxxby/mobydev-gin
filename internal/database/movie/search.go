package movie

import (
	"errors"
	"project/internal/utils"
)

// searches moive by query looking at the name, or keywords order by watch count
func (db *MovieRepository) SearchMovie(userId int, query string) ([]Movie, error) {
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
		ORDER BY m.watch_count DESC
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
		genreRows, err := db.Database.Query("SELECT g.name FROM genres g INNER JOIN movie_genres mg ON g.id = mg.genre_id WHERE mg.movie_id = $1", movie.Id)
		if err != nil {
			return nil, err
		}
		defer genreRows.Close()
		var countSeasons int
		var seasonId int
		var countEpisodes int

		db.Database.QueryRow("SELECT COUNT(*) FROM seasons WHERE movie_id = $1", movie.Id).Scan(&countSeasons)
		db.Database.QueryRow("SELECT id FROM seasons WHERE movie_id = $1;", movie.Id).Scan(&seasonId)
		db.Database.QueryRow("SELECT COUNT(*) FROM episodes WHERE season_id = $1", seasonId).Scan(&countEpisodes)
		movie.SeasonCount = countSeasons
		movie.SeriesCount = countEpisodes
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

// RIGHT NOW DONT WORKS an gives only 5 random movies form db
func (db *MovieRepository) GetSimularMoviesLimit5(excludeMovie int) ([]Movie, error) {
	var count int
	err := db.Database.QueryRow("SELECT COUNT(*) FROM movies").Scan(&count)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.New("no movies found in the database")
	}

	randomIndexes := utils.GenerateRandomIndexes(count, 5)

	rows, err := db.Database.Query(`
		SELECT m.id, m.name, p.main_poster
		FROM movies m
		LEFT JOIN posters p ON m.id = p.movie_id
		WHERE m.id != $2
		LIMIT 5 OFFSET $1 
	`, randomIndexes[0], excludeMovie)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		rows.Scan(&movie.Id, &movie.Name, &movie.Poster[0])
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}
