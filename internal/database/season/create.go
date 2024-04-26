package season

import (
	"time"
)

func (DB *SeasonRepository) AddSeasonToTheMovie(movieId, userId int, name string, seasonNumber int, desc string, releaseDate time.Time) (int, error) {
	tx, err := DB.Database.Begin()
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()
	q := "INSERT INTO seasons  (user_id, movie_id, season_number , name , description , release_date) VALUES ($1,$2,$3,$4,$5,$6)RETURNING id "
	var seasonId int
	err = tx.QueryRow(q, userId, movieId, seasonNumber, name, desc, releaseDate).Scan(&seasonId)
	if err != nil {
		return -1, err
	}
	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return seasonId, nil
}
