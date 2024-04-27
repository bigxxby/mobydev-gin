package episode

func (db *EpisodeRepository) DeleteEpisodeById(episodeId int) error {
	tx, err := db.Database.Begin()
	if err != nil {

		return err
	}
	defer tx.Rollback()
	q := "DELETE FROM episodes WHERE id = $1"
	_, err = tx.Exec(q, episodeId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
func (db *EpisodeRepository) DeleteAllEpisodesByIdOfSeason(seasonId int) error {
	tx, err := db.Database.Begin()
	if err != nil {

		return err
	}
	defer tx.Rollback()
	q := "DELETE FROM episodes WHERE season_id = $1"
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
func (db *EpisodeRepository) DeleteEpisodeByNumberOfSelectedSeason(seasonId int, episodeNumber int) error {
	tx, err := db.Database.Begin()
	if err != nil {

		return err
	}
	defer tx.Rollback()
	q := "DELETE FROM episodes WHERE season_id = $1 AND episode_number = $2"
	_, err = tx.Exec(q, seasonId, episodeNumber)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
