package database

import "log"

func (db *Database) GetUserById(id int) (*User, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var user User
	err := db.Database.QueryRow("SELECT id , email , name , phone , date_of_birth , role FROM users WHERE id = $1", id).Scan(&user.Id, &user.Email, &user.Name, &user.Phone, &user.DateOfBirth, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (db *Database) GetMovies(limit int) ([]Movie, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if limit != 0 {

		stmt, err := db.Database.Prepare("SELECT * FROM movies ORDER BY created_at LIMIT $1 DESC")
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
				&movie.Category,
				&movie.MovieType,
				&movie.Year,
				&movie.AgeCategory,
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

		stmt, err := db.Database.Prepare("SELECT * FROM movies ORDER BY  created_at DESC")
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
				&movie.Category,
				&movie.MovieType,
				&movie.Year,
				&movie.AgeCategory,
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
func (db *Database) GetMovieById(id int) (*Movie, error) {
	var movie Movie
	err := db.Database.QueryRow(
		"SELECT * FROM movies WHERE id = $1", id,
	).Scan(
		&movie.Id,
		&movie.UserId,
		&movie.ImageUrl,
		&movie.Name,
		&movie.Category,
		&movie.MovieType,
		&movie.Year,
		&movie.AgeCategory,
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

func (db *Database) GetSeasonById(id int) (*Season, error) {
	var season Season

	query := `SELECT id, user_id, movie_id, season_number, name, description, release_date FROM seasons WHERE id = $1`

	err := db.Database.QueryRow(query, id).Scan(
		&season.ID,
		&season.UserID,
		&season.MovieID,
		&season.SeasonNumber,
		&season.Name,
		&season.Description,
		&season.ReleaseDate,
	)

	if err != nil {
		return nil, err
	}

	return &season, nil
}
func (db *Database) GetEpisodeById(id int) (*Episode, error) {
	var episode Episode

	query := `SELECT id, user_id, url, season_id, episode_number, name, duration_minutes, release_date, description 
	FROM episodes WHERE id = $1`

	err := db.Database.QueryRow(query, id).Scan(
		&episode.ID,
		&episode.UserID,
		&episode.URL,
		&episode.SeasonID,
		&episode.EpisodeNumber,
		&episode.Name,
		&episode.DurationMinutes,
		&episode.ReleaseDate,
		&episode.Description,
	)

	if err != nil {
		return nil, err
	}

	return &episode, nil
}
func (db *Database) GetTrendsById(id int) (*Trend, error) {
	var trend Trend

	query := `
	SELECT id, movies_id, trend_date, trend_value
	FROM trends
	WHERE id = $1
	`

	err := db.Database.QueryRow(query, id).Scan(
		&trend.ID,
		&trend.MovieID,
		&trend.TrendDate,
		&trend.TrendValue,
	)

	if err != nil {
		return nil, err
	}

	return &trend, nil
}
func (db *Database) GetTrends() ([]*Trend, error) {
	var trends []*Trend

	query := `
	SELECT *
	FROM trends
	ORDER BY trend_date DESC
	`

	rows, err := db.Database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var trend Trend
		err := rows.Scan(
			&trend.ID,
			&trend.MovieID,
			&trend.TrendDate,
			&trend.TrendValue,
		)
		if err != nil {
			return nil, err
		}
		trends = append(trends, &trend)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return trends, nil
}
