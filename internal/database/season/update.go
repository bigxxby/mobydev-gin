package season

import "time"

func (db *SeasonRepository) UpdateSeason(seasonId int, seasonNumber int, name string, desc string, releaseDate time.Time) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := `
        UPDATE seasons 
        SET season_number = $1, 
            name = $2, 
            description = $3, 
            release_date = $4
        WHERE id = $5 
    `

	_, err = tx.Exec(q, seasonNumber, name, desc, releaseDate, seasonId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
