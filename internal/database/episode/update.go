package episode

import "time"

func (db *EpisodeRepository) UpdateEpisode(episodeID int, userID int, url string, episodeNumber int, name string, durationMinutes int, releaseDate time.Time, description string) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := `
        UPDATE episodes 
        SET user_id = $1, 
            episode_number = $2, 
            name = $3, 
            duration_minutes = $4, 
            release_date = $5, 
            description = $6
        WHERE id = $7
    `

	_, err = tx.Exec(q, userID, episodeNumber, name, durationMinutes, releaseDate, description, episodeID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
