package trend

import (
	"database/sql"
	"time"
)

type TrendRepository struct {
	Database *sql.DB
}
type Trend struct {
	ID         int       `json:"id"`
	MovieID    int       `json:"movie_id"`
	TrendDate  time.Time `json:"trend_date"`
	TrendValue int       `json:"trend_value"`
}
