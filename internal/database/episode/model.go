package episode

import (
	"database/sql"
	"time"
)

type EpisodeRepository struct {
	Database *sql.DB
}
type Episode struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	URL             string    `json:"url"`
	SeasonID        int       `json:"season_id"`
	EpisodeNumber   int       `json:"episode_number"`
	Name            string    `json:"name"`
	DurationMinutes int       `json:"duration_minutes"`
	ReleaseDate     time.Time `json:"release_date"`
	Description     string    `json:"description"`
}
