package movie

import "log"

func (db *MovieRepository) CreateMovie(userId int, imageUrl string, name string, year int, categoryId, ageCategoryId, genreId int, durationMinutes int, keywords string, desc string, director string, producer string) (*Movie, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	tx, err := db.Database.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO movies (
			user_id, image_url, name, year, category_id, age_category_id, genre_id, duration_minutes, keywords, description, director, producer
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id, created_at, updated_at
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement
	var movie Movie
	err = stmt.QueryRow(
		userId, imageUrl, name, year, categoryId, ageCategoryId, genreId, durationMinutes, keywords, desc, director, producer,
	).Scan(
		&movie.Id, &movie.CreatedAt, &movie.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &movie, nil
}
