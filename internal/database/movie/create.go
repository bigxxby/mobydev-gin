package movie

import "log"

func (db *MovieRepository) CreateMovie(userId int, imageUrl string, name string, category string, movieType string, year int, ageCategory string, durationMinutes int, keywords string, desc string, director string, producer string) (*Movie, error) {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Start a new transaction
	tx, err := db.Database.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Prepare the SQL statement for inserting a new movie
	stmt, err := tx.Prepare(`
		INSERT INTO movies (
			user_id, image_url, name, category, movie_type, year, age_category, duration_minutes, keywords, description, director, producer
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id, created_at, updated_at
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement
	var movie Movie
	err = stmt.QueryRow(
		userId, imageUrl, name, category, movieType, year, ageCategory, durationMinutes, keywords, desc, director, producer,
	).Scan(
		&movie.Id, &movie.CreatedAt, &movie.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &movie, nil
}
