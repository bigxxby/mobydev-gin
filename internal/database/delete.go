package database

import "log"

func (db *Database) DeleteMovie(movieId string) error {
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

	// Удаляем связанные записи из таблицы "trends" перед удалением фильма
	_, err = tx.Exec("DELETE FROM trends WHERE movie_id=$1", movieId)
	if err != nil {
		return err
	}

	// Удаляем эпизоды, связанные с фильмом
	_, err = tx.Exec("DELETE FROM episodes WHERE season_id IN (SELECT id FROM seasons WHERE movie_id=$1)", movieId)
	if err != nil {
		return err
	}

	// Удаляем сезоны, связанные с фильмом
	_, err = tx.Exec("DELETE FROM seasons WHERE movie_id=$1", movieId)
	if err != nil {
		return err
	}

	// Удаляем фильм
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
func (db *Database) DeleteAllFavoritesByUserId(userId int) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM favorites WHERE user_id = $1", userId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
