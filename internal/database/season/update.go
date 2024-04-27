package season

import "time"

func (db *SeasonRepository) UpdateSeason(seasonId int, userId int, seasonNumber int, name string, desc string, releaseDate time.Time) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := `
        UPDATE seasons 
        SET user_id = $1, 
            season_number = $2, 
            name = $3, 
            description = $4, 
            release_date = $5
        WHERE id = $6 
    `

	_, err = tx.Exec(q, userId, seasonNumber, name, desc, releaseDate, seasonId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
