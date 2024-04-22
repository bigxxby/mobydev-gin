package movie

import (
	"strings"
)

func (db *MovieRepository) SearchMovie(query string) ([]Movie, error) {
	query = "%" + query + "%"
	query = strings.ToLower(query)

	sqlQuery := `
	SELECT * FROM movies 
		WHERE to_tsvector('english', name || ' ' || keywords) @@ plainto_tsquery('english', $1)
	`

	rows, err := db.Database.Query(sqlQuery, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie

	for rows.Next() {
		var movie Movie
		err := rows.Scan(
			&movie.Id,
			&movie.UserId,
			&movie.ImageUrl,
			&movie.Name,
			&movie.Year,
			&movie.CategoryId,
			&movie.AgeCategoryId,
			&movie.GenreId,
			&movie.DurationMinutes,
			&movie.Keywords,
			&movie.Description,
			&movie.Director,
			&movie.Producer,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

// func (db *MovieRepository) SearchMovieByGenre(genre string) ([]Movie, error) {
// 	genre = strings.ToLower(genre)
// 	q := `
// 	SELECT * FROM movies m INNER JOIN genres g ON m.genre_id = g.id WHERE LOWER(g.name) = $1
// 	`

// 	rows, err := db.Database.Query(q, genre)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var movies []Movie

// 	for rows.Next() {
// 		var movie Movie
// 		err := rows.Scan(
// 			&movie.Id,
// 			&movie.UserId,
// 			&movie.ImageUrl,
// 			&movie.Name,
// 			&movie.Year,
// 			&movie.CategoryId,
// 			&movie.AgeCategoryId,
// 			&movie.GenreId,
// 			&movie.DurationMinutes,
// 			&movie.Keywords,
// 			&movie.Description,
// 			&movie.Director,
// 			&movie.Producer,
// 			&movie.CreatedAt,
// 			&movie.UpdatedAt,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		movies = append(movies, movie)
// 	}

// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return movies, nil
// }
