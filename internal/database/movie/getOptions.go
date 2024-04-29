package movie

func (db *MovieRepository) GetEveryMovieByGenre() (map[string][]MovieShort, error) {

	moviesByGenre := make(map[string][]MovieShort)

	query := `
        SELECT g.name, m.name, m.id , p.main_poster
        FROM movies m
        INNER JOIN movie_genres mg ON m.id = mg.movie_id
        INNER JOIN genres g ON mg.genre_id = g.id
        LEFT JOIN posters p ON m.id = p.movie_id
    `

	rows, err := db.Database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var genre string
		var movie MovieShort
		rows.Scan(&genre, &movie.Name, &movie.Id, &movie.MainPoster)

		moviesByGenre[genre] = append(moviesByGenre[genre], movie)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return moviesByGenre, nil
}
func (db *MovieRepository) GetEveryMovieByCategory() (map[string][]MovieShort, error) {

	moviesByCategory := make(map[string][]MovieShort)

	query := `
        SELECT c.name, m.name, m.id , p.main_poster
        FROM movies m
        INNER JOIN categories c ON m.category_id = c.id
        LEFT JOIN posters p ON m.id = p.movie_id
    `

	rows, err := db.Database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category string
		var movie MovieShort
		rows.Scan(&category, &movie.Name, &movie.Id, &movie.MainPoster)

		moviesByCategory[category] = append(moviesByCategory[category], movie)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return moviesByCategory, nil
}
func (db *MovieRepository) GetEveryMovieByAgeCategory() (map[string][]MovieShort, error) {

	moviesByAgeCategory := make(map[string][]MovieShort)

	query := `
        SELECT ac.name, m.name, m.id , p.main_poster
        FROM movies m
        INNER JOIN age_categories ac ON m.age_category_id = ac.id
        LEFT JOIN posters p ON m.id = p.movie_id
    `

	rows, err := db.Database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var age string
		var movie MovieShort
		rows.Scan(&age, &movie.Name, &movie.Id, &movie.MainPoster)

		moviesByAgeCategory[age] = append(moviesByAgeCategory[age], movie)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return moviesByAgeCategory, nil
}
