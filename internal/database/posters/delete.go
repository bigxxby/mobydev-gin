package posters

func (db *PosterRepo) DeletePosters(posterId int) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := "DELETE FROM posters WHERE id = $1"

	_, err = tx.Exec(query, posterId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
func (db *PosterRepo) DeletePostersOfMovie(movieId int) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := "DELETE FROM posters WHERE movie_id = $1"

	_, err = tx.Exec(query, movieId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
