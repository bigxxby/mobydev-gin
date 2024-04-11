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

	// Deleting related records from the "episodes" table
	_, err = tx.Exec("DELETE FROM episodes WHERE season_id IN (SELECT id FROM seasons WHERE movie_id=$1)", movieId)
	if err != nil {
		return err
	}

	// Deleting related records from the "seasons" table
	_, err = tx.Exec("DELETE FROM seasons WHERE movie_id=$1", movieId)
	if err != nil {
		return err
	}

	// Deleting the record from the "movies" table
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

// func (db *Database) DeleteUser(userID string) error {
// 	log.SetFlags(log.LstdFlags | log.Lshortfile)

// 	tx, err := db.Database.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	defer tx.Rollback()

// 	_, err = tx.Exec("DELETE FROM users WHERE id=$1", userID)
// 	if err != nil {
// 		return err
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
