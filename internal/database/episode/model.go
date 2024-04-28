package episode

import (
	"database/sql"
)

type EpisodeRepository struct {
	Database *sql.DB
}

type Episode struct {
	ID              int    `json:"id"`
	UserID          int    `json:"user_id"`
	URL             string `json:"url" binding:"required"`
	SeasonID        int    `json:"season_id"`
	EpisodeNumber   int    `json:"episode_number" binding:"required"`
	Name            string `json:"name" binding:"required"`
	DurationMinutes int    `json:"duration_minutes" binding:"required"`
	ReleaseDate     string `json:"release_date" binding:"required"`
	Description     string `json:"description"`
}
