package genres

import "database/sql"

func (db *GenreRepository) GetAllGenres() ([]Genre, error) {
	query := "SELECT * FROM genres"

	rows, err := db.Database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []Genre

	for rows.Next() {
		var genre Genre
		if err := rows.Scan(&genre.ID, &genre.UserID, &genre.Name, &genre.Description); err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return genres, nil
}
func (db *GenreRepository) GetGenreById(id int) (*Genre, error) {
	query := "SELECT * FROM genres WHERE id = $1"

	row := db.Database.QueryRow(query, id)

	var genre Genre

	err := row.Scan(&genre.ID, &genre.UserID, &genre.Name, &genre.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &genre, nil
}
