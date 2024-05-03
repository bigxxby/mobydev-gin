package genres

import "database/sql"

type GenreRepository struct {
	Database *sql.DB
}
type Genre struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description" binding:"required"`
}
type GenreShort struct {
	Name        string `json:"name"`
	Description string `json:"description" binding:"required"`
}
