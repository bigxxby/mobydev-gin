package favorites

import (
	"database/sql"
	"project/internal/database/movie"
	"time"
)

type FavoritesRepository struct {
	Database *sql.DB
}
type Favorite struct {
	ID      int         `json:"id"`
	UserID  int         `json:"user_id"`
	MovieID int         `json:"movie_id"`
	AddedAt time.Time   `json:"added_at"`
	Movie   movie.Movie `json:"movie"`
}
