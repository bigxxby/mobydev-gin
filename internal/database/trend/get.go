package trend

func (db *TrendRepository) GetTrendsById(id int) (*Trend, error) {
	var trend Trend

	query := `
	SELECT id, movie_id, trend_date, trend_value
	FROM trends
	WHERE id = $1
	`

	err := db.Database.QueryRow(query, id).Scan(
		&trend.ID,
		&trend.MovieID,
		&trend.TrendDate,
		&trend.TrendValue,
	)

	if err != nil {
		return nil, err
	}

	return &trend, nil
}
func (db *TrendRepository) GetTrends() ([]*Trend, error) {
	var trends []*Trend

	query := `
	SELECT *
	FROM trends
	ORDER BY trend_date DESC
	`

	rows, err := db.Database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var trend Trend
		err := rows.Scan(
			&trend.ID,
			&trend.MovieID,
			&trend.TrendDate,
			&trend.TrendValue,
		)
		if err != nil {
			return nil, err
		}
		trends = append(trends, &trend)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return trends, nil
}
