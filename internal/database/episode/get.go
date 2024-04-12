package episode

func (db *EpisodeRepository) GetEpisodeById(id int) (*Episode, error) {
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
