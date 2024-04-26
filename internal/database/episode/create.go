package episode

import "time"

func (db *EpisodeRepository) CreateEpisode(userID, seasonID, episodeNumber int, url, name, description string, durationMinutes int, releaseDate time.Time) (int, error) {
	var episodeID int

	query := `
        INSERT INTO episodes (user_id, season_id, episode_number, url, name, description, duration_minutes, release_date)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id;
    `

	err := db.Database.QueryRow(query, userID, seasonID, episodeNumber, url, name, description, durationMinutes, releaseDate).Scan(&episodeID)
	if err != nil {
		return 0, err
	}

	return episodeID, nil
}
