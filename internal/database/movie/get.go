package movie

import (
	"log"
)

func (db *MovieRepository) GetMovieById(id int) (*Movie, error) {
	var movie Movie
	err := db.Database.QueryRow(
		"SELECT * FROM movies WHERE id = $1", id,
	).Scan(
		&movie.Id,
		&movie.UserId,
		&movie.ImageUrl,
		&movie.Name,
		&movie.Year,
		&movie.CategoryId,
		&movie.AgeCategoryId,
		&movie.GenreId,
		&movie.WatchCount,
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
	return &movie, nil
}
func (db *MovieRepository) GetMovies(limit int) ([]Movie, error) {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if limit != 0 {

		stmt, err := db.Database.Prepare("SELECT * FROM movies LIMIT $1 ")
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
			err := rows.Scan(
				&movie.Id,
				&movie.UserId,
				&movie.ImageUrl,
				&movie.Name,
				&movie.Year,
				&movie.CategoryId,
				&movie.AgeCategoryId,
				&movie.GenreId,
				&movie.WatchCount,
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

	} else {

		stmt, err := db.Database.Prepare("SELECT * FROM movies ")
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
			err := rows.Scan(
				&movie.Id,
				&movie.UserId,
				&movie.ImageUrl,
				&movie.Name,
				&movie.Year,
				&movie.CategoryId,
				&movie.AgeCategoryId,
				&movie.GenreId,
				&movie.WatchCount,
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
}
