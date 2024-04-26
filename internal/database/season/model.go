package season

import (
	"database/sql"
	"project/internal/database/episode"
)

type SeasonRepository struct {
	Database *sql.DB
}
type Season struct {
	ID           int               `json:"id"`
	UserID       int               `json:"user_id"`
	MovieID      int               `json:"movie_id"`
	SeasonNumber int               `json:"season_number" binding:"required"`
	Name         string            `json:"name" binding:"required"`
	Description  string            `json:"description" binding:"required"`
	ReleaseDate  string            `json:"release_date" binding:"required"`
	Episodes     []episode.Episode `json:"episodes"`
}
