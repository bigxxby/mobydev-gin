package season

func (db *SeasonRepository) GetSeasonById(id int) (*Season, error) {
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
