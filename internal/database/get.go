package database

import "log"

func (db *Database) GetUserById(id int) (*User, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	tx, err := db.Database.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	var user User
	err = tx.QueryRow("SELECT id , email , name , phone , date_of_birth , role FROM users WHERE id = $1", id).Scan(&user.Id, &user.Email, &user.Name, &user.Phone, &user.DateOfBirth, &user.Role)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (db *Database) GetProjects(limit int) ([]Project, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if limit != 0 {
		tx, err := db.Database.Begin()
		if err != nil {
			return nil, err
		}
		defer tx.Rollback()

		stmt, err := tx.Prepare("SELECT * FROM projects  ORDER BY created_at LIMIT $1")
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		rows, err := stmt.Query(limit)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var projects []Project

		for rows.Next() {
			var project Project
			err := rows.Scan(
				&project.Id,
				&project.UserId,
				&project.ImageUrl,
				&project.Name,
				&project.Category,
				&project.ProjectType,
				&project.Year,
				&project.AgeCategory,
				&project.DurationMinutes,
				&project.Keywords,
				&project.Description,
				&project.Director,
				&project.Producer,
				&project.CreatedAt,
				&project.UpdatedAt,
			)
			if err != nil {
				return nil, err
			}
			projects = append(projects, project)
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}

		err = tx.Commit()
		if err != nil {
			return nil, err
		}

		return projects, nil

	} else {

		tx, err := db.Database.Begin()
		if err != nil {
			return nil, err
		}
		defer tx.Rollback()

		stmt, err := tx.Prepare("SELECT * FROM projects ORDER BY created_at")
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		rows, err := stmt.Query()
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var projects []Project

		for rows.Next() {
			var project Project
			err := rows.Scan(
				&project.Id,
				&project.UserId,
				&project.ImageUrl,
				&project.Name,
				&project.Category,
				&project.ProjectType,
				&project.Year,
				&project.AgeCategory,
				&project.DurationMinutes,
				&project.Keywords,
				&project.Description,
				&project.Director,
				&project.Producer,
				&project.CreatedAt,
				&project.UpdatedAt,
			)
			if err != nil {
				return nil, err
			}
			projects = append(projects, project)
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}

		err = tx.Commit()
		if err != nil {
			return nil, err
		}

		return projects, nil
	}
}
func (db *Database) GetProjectById(id int) (*Project, error) {
	var project Project
	err := db.Database.QueryRow(
		"SELECT * FROM projects WHERE id = $1", id,
	).Scan(
		&project.Id,
		&project.UserId,
		&project.ImageUrl,
		&project.Name,
		&project.Category,
		&project.ProjectType,
		&project.Year,
		&project.AgeCategory,
		&project.DurationMinutes,
		&project.Keywords,
		&project.Description,
		&project.Director,
		&project.Producer,
		&project.CreatedAt,
		&project.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (db *Database) GetSeasonById(id int) (*Season, error) {
	var season Season

	query := `SELECT id, user_id, project_id, season_number, name, description, release_date FROM seasons WHERE id = $1`

	err := db.Database.QueryRow(query, id).Scan(
		&season.ID,
		&season.UserID,
		&season.ProjectID,
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
	SELECT id, project_id, trend_date, trend_value
	FROM trends
	WHERE id = $1
	`

	err := db.Database.QueryRow(query, id).Scan(
		&trend.ID,
		&trend.ProjectID,
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
			&trend.ProjectID,
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
