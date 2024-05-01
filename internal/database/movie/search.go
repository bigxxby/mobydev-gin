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

// func (db *MovieRepository) GetMoviesByFilter(userId int, option string, optionData string) ([]*Movie, error) {
// 	var movies []*Movie
// 	var err error

// 	switch option {
// 	case "genre":
// 		movies, err = db.GetMoviesByGenre(userId, optionData)
// 	case "category":
// 		movies, err = db.GetMoviesByCategory(userId, optionData)
// 	case "ageCategory":
// 		movies, err = db.GetMoviesByAgeCategory(userId, optionData)
// 	default:
// 		return nil, errors.New("unsupported option")
// 	}

// 	if err != nil {
// 		return nil, err
// 	}

// 	return movies, nil
// }

// func (db *MovieRepository) GetMoviesByGenre(userId int, genre string) ([]*Movie, error) {
// 	return db.getMoviesByCategoryQuery(userId, genre, "g.name = $1")
// }

// func (db *MovieRepository) GetMoviesByAgeCategory(userId int, ageCategory string) ([]*Movie, error) {
// 	return db.getMoviesByCategoryQuery(userId, ageCategory, "ac.name = $1")
// }

// func (db *MovieRepository) GetMoviesByCategory(userId int, category string) ([]*Movie, error) {
// 	return db.getMoviesByCategoryQuery(userId, category, "c.name = $1")

// }

// func (db *MovieRepository) getMoviesByCategoryQuery(userId int, filterValue string, filterCondition string) ([]*Movie, error) {
// 	query := `
//         SELECT m.id, m.user_id, m.name, m.year, m.category_id, m.age_category_id,
//                m.watch_count, m.duration_minutes, m.keywords, m.description,
//                m.director, m.producer, m.created_at, m.updated_at,
//                c.name AS category_name, ac.name AS age_category_name,
//                p.main_poster, p.second_poster, p.third_poster, p.fourth_poster, p.fifth_poster
//         FROM movies m
//         LEFT JOIN categories c ON m.category_id = c.id
//         LEFT JOIN age_categories ac ON m.age_category_id = ac.id
//         LEFT JOIN posters p ON m.id = p.movie_id
//         INNER JOIN movie_genres mg ON m.id = mg.movie_id
//         INNER JOIN genres g ON mg.genre_id = g.id
//         WHERE ` + filterCondition + `
//         ORDER BY m.watch_count DESC
//     `

// 	rows, err := db.Database.Query(query, filterValue)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var movies []*Movie
// 	for rows.Next() {
// 		var movie Movie
// 		rows.Scan(&movie.Id, &movie.UserId, &movie.Name, &movie.Year,
// 			&movie.CategoryId, &movie.AgeCategoryId, &movie.WatchCount, &movie.DurationMinutes,
// 			&movie.Keywords, &movie.Description, &movie.Director, &movie.Producer,
// 			&movie.CreatedAt, &movie.UpdatedAt, &movie.Category, &movie.AgeCategory,
// 			&movie.Poster[0], &movie.Poster[1], &movie.Poster[2],
// 			&movie.Poster[3], &movie.Poster[4])
// 		if err != nil {
// 			return nil, err
// 		}

// 		var isFavorite int
// 		err = db.Database.QueryRow("SELECT COUNT(*) FROM favorites WHERE user_id = $1 AND movie_id = $2", userId, movie.Id).Scan(&isFavorite)
// 		if err != nil {
// 			return nil, err
// 		}
// 		movie.IsFavorite = isFavorite > 0

// 		genreQuery := `
//             SELECT g.name
//             FROM genres g
//             INNER JOIN movie_genres mg ON g.id = mg.genre_id
//             WHERE mg.movie_id = $1
//         `
// 		genreRows, err := db.Database.Query(genreQuery, movie.Id)
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer genreRows.Close()
// 		var countSeasons int
// 		var seasonId int
// 		var countEpisodes int
// 		db.Database.QueryRow("SELECT COUNT(*) FROM seasons WHERE movie_id = $1", movie.Id).Scan(&countSeasons)
// 		db.Database.QueryRow("SELECT id FROM seasons WHERE movie_id = $1;", movie.Id).Scan(&seasonId)
// 		db.Database.QueryRow("SELECT COUNT(*) FROM episodes WHERE season_id = $1", seasonId).Scan(&countEpisodes)
// 		movie.SeasonCount = countSeasons
// 		movie.SeriesCount = countEpisodes

// 		var genresArr []string
// 		for genreRows.Next() {
// 			var genre string
// 			err := genreRows.Scan(&genre)
// 			if err != nil {
// 				return nil, err
// 			}
// 			genresArr = append(genresArr, genre)
// 		}
// 		movie.Genres = genresArr

// 		movies = append(movies, &movie)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return movies, nil
// }

// func (db *MovieRepository) GetMoviesByGenre(userId int, genre string) ([]*Movie, error) {
// 	query := `
//         SELECT m.id, m.user_id, m.name, m.year, m.category_id, m.age_category_id,
//                m.watch_count, m.duration_minutes, m.keywords, m.description,
//                m.director, m.producer, m.created_at, m.updated_at,
//                c.name AS category_name, ac.name AS age_category_name,
//                p.main_poster, p.second_poster, p.third_poster, p.fourth_poster, p.fifth_poster
//         FROM movies m
//         LEFT JOIN categories c ON m.category_id = c.id
//         LEFT JOIN age_categories ac ON m.age_category_id = ac.id
//         LEFT JOIN posters p ON m.id = p.movie_id
//         INNER JOIN movie_genres mg ON m.id = mg.movie_id
//         INNER JOIN genres g ON mg.genre_id = g.id
//         WHERE g.name = $1 ORDER BY m.watch_count DESC
//     `

// 	rows, err := db.Database.Query(query, genre)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var movies []*Movie
// 	for rows.Next() {
// 		var movie Movie
// 		err := rows.Scan(&movie.Id, &movie.UserId, &movie.Name, &movie.Year,
// 			&movie.CategoryId, &movie.AgeCategoryId, &movie.WatchCount, &movie.DurationMinutes,
// 			&movie.Keywords, &movie.Description, &movie.Director, &movie.Producer,
// 			&movie.CreatedAt, &movie.UpdatedAt, &movie.Category, &movie.AgeCategory,
// 			&movie.Poster[0], &movie.Poster[1], &movie.Poster[2],
// 			&movie.Poster[3], &movie.Poster[4])
// 		if err != nil {
// 			return nil, err
// 		}

// 		var isFavorite int
// 		err = db.Database.QueryRow("SELECT COUNT(*) FROM favorites WHERE user_id = $1 AND movie_id = $2", userId, movie.Id).Scan(&isFavorite)
// 		if err != nil {
// 			return nil, err
// 		}
// 		movie.IsFavorite = isFavorite > 0

// 		genreQuery := `
//             SELECT g.name
//             FROM genres g
//             INNER JOIN movie_genres mg ON g.id = mg.genre_id
//             WHERE mg.movie_id = $1
//         `
// 		genreRows, err := db.Database.Query(genreQuery, movie.Id)
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer genreRows.Close()

// 		var genresArr []string
// 		for genreRows.Next() {
// 			var genre string
// 			err := genreRows.Scan(&genre)
// 			if err != nil {
// 				return nil, err
// 			}
// 			genresArr = append(genresArr, genre)
// 		}
// 		movie.Genres = genresArr

//			movies = append(movies, &movie)
//		}
//		if err := rows.Err(); err != nil {
//			return nil, err
//		}
//		return movies, nil
//	}
