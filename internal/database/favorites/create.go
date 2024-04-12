package favorites

import "time"

func (db *FavoritesRepository) AddToFavorites(userId, movieId int) (*Favorite, error) {
	tx, err := db.Database.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `
    INSERT INTO favorites (user_id, movie_id, added_at) 
    VALUES ($1, $2, CURRENT_TIMESTAMP)
    RETURNING id, added_at
    `

	var favoriteId int
	var addedAt time.Time
	err = tx.QueryRow(query, userId, movieId).Scan(&favoriteId, &addedAt)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	favorite := &Favorite{
		ID:      favoriteId,
		UserID:  userId,
		MovieID: movieId,
		AddedAt: addedAt,
	}

	return favorite, nil
}
