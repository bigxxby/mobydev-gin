package season

import (
	"database/sql"
	"time"
)

type SeasonRepository struct {
	Database *sql.DB
}
type Season struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	MovieID      int       `json:"movie_id"`
	SeasonNumber int       `json:"season_number"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	ReleaseDate  time.Time `json:"release_date"`
}
