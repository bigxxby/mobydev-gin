package favorites

func (db *FavoritesRepository) CheckIfMovieAddedToFavorites(userId, movieId int) (bool, error) {
	var exists bool

	err := db.Database.QueryRow("SELECT EXISTS(SELECT 1 FROM favorites WHERE user_id=$1 AND movie_id=$2)", userId, movieId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
