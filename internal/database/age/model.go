package age

import "database/sql"

type AgeRepository struct {
	Database *sql.DB
}

type AgeCategory struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `json:"name" binding:"required"`
	Note   string `json:"note"`
	MinAge int    `json:"min_age" binding:"required"`
	MaxAge int    `json:"max_age" binding:"required"`
}
