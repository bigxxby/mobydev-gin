package movie

import "log"

func (db *MovieRepository) DeleteMovie(movieId string) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.Exec("DELETE FROM favorites WHERE movie_id=$1", movieId)
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM trends WHERE movie_id=$1", movieId)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM episodes WHERE season_id IN (SELECT id FROM seasons WHERE movie_id=$1)", movieId)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM seasons WHERE movie_id=$1", movieId)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM movies WHERE id=$1", movieId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
