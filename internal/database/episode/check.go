package episode

import "database/sql"

func (db *EpisodeRepository) CheckEpisodeExistsById(episodeId int) error {

	stmt, err := db.Database.Prepare("SELECT * FROM episodes WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(episodeId)
	if err != nil {
		return err
	}
	defer rows.Close()

	exists := rows.Next()
	if !exists {
		return sql.ErrNoRows
	}

	return nil
}
