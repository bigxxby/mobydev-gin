package posters

func (db *PosterRepo) GetPostersAllOfMovieById(movieId int) (*Poster, error) {
	query := `
		SELECT * FROM posters WHERE movie_id = $1 LIMIT 1
	`
	var poster Poster

	row := db.Database.QueryRow(query, movieId)

	err := row.Scan(&poster.Id, &poster.MovieId, &poster.MainPoster, &poster.SecondaryPoster, &poster.ThirdPoster, &poster.FourthPoster, &poster.FifthPoster)
	if err != nil {
		return nil, err
	}
	return &poster, nil
}
