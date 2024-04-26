package season

import "project/internal/database/episode"

func (db *SeasonRepository) GetSeasonById(id int) (*Season, error) {
	var season Season

	// Запрос для получения информации о сезоне
	seasonQuery := `
        SELECT id, user_id, movie_id, season_number, name, description, release_date 
        FROM seasons WHERE id = $1
    `

	// Выполняем запрос и сканируем данные в объект season
	err := db.Database.QueryRow(seasonQuery, id).Scan(
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

	// Запрос для получения всех эпизодов данного сезона
	episodesQuery := `
        SELECT id, user_id, url, season_id, episode_number, name, duration_minutes, release_date, description 
        FROM episodes WHERE season_id = $1
    `

	// Выполняем запрос и получаем все эпизоды для данного сезона
	rows, err := db.Database.Query(episodesQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Проходим по всем эпизодам и добавляем их в объект season
	for rows.Next() {
		var episode episode.Episode
		err := rows.Scan(
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
		// Добавляем эпизод в массив эпизодов сезона
		season.Episodes = append(season.Episodes, episode)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &season, nil
}

func (db *SeasonRepository) GetAllSeasonsOfMovieId(movieID int) ([]Season, error) {
	var seasons []Season

	querySeasons := `SELECT * FROM seasons WHERE movie_id = $1 ORDER BY season_number`

	rowsSeasons, err := db.Database.Query(querySeasons, movieID)
	if err != nil {
		return nil, err
	}
	defer rowsSeasons.Close()

	for rowsSeasons.Next() {
		var season Season
		err := rowsSeasons.Scan(
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

		queryEpisodes := `SELECT * FROM episodes WHERE season_id = $1 ORDER BY episode_number`

		rowsEpisodes, err := db.Database.Query(queryEpisodes, season.ID)
		if err != nil {
			return nil, err
		}
		defer rowsEpisodes.Close()

		for rowsEpisodes.Next() {
			var episode episode.Episode
			err := rowsEpisodes.Scan(
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
			season.Episodes = append(season.Episodes, episode)
		}

		if err := rowsEpisodes.Err(); err != nil {
			return nil, err
		}

		seasons = append(seasons, season)
	}

	if err := rowsSeasons.Err(); err != nil {
		return nil, err
	}

	return seasons, nil
}
