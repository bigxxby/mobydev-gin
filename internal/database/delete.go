package database

import "log"

func (db *Database) DeleteProject(projectId string) error {
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

	// Удаление связанных записей из таблицы "episodes"
	_, err = tx.Exec("DELETE FROM episodes WHERE season_id IN (SELECT id FROM seasons WHERE project_id=$1)", projectId)
	if err != nil {
		return err
	}

	// Удаление связанных записей из таблицы "seasons"
	_, err = tx.Exec("DELETE FROM seasons WHERE project_id=$1", projectId)
	if err != nil {
		return err
	}

	// Удаление записи из таблицы "projects"
	_, err = tx.Exec("DELETE FROM projects WHERE id=$1", projectId)
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
