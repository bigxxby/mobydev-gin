package season

func (db *SeasonRepository) DeleteSeason(seasonId int) error {
	tx, err := db.Database.Begin()
	if err != nil {

		return err
	}
	defer tx.Rollback()
	q := "DELETE FROM seasons WHERE id = $1"
	_, err = tx.Exec(q, seasonId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {

		return err
	}
	return nil
}
func (db *SeasonRepository) DeleteSeasonNumberOfSelectedMovie(movieId int, seasonNumber int) error {
	tx, err := db.Database.Begin()
	if err != nil {

		return err
	}
	defer tx.Rollback()
	q := "DELETE FROM seasons WHERE movie_id = $1 AND season_number = $2"
	_, err = tx.Exec(q, movieId, seasonNumber)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (db *SeasonRepository) DeleteAllSeasonsOfCurrentMovie(movieId int) error {
	tx, err := db.Database.Begin()
	if err != nil {

		return err
	}
	defer tx.Rollback()
	q := "DELETE FROM seasons WHERE movie_id = $1"
	_, err = tx.Exec(q, movieId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
