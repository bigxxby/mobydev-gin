package categories

import (
	"database/sql"
)

type CategoryRepository struct {
	Database *sql.DB
}
type Category struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id" ` // created by
	Name        string `json:"category_name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
