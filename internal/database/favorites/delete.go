package favorites

func (db *FavoritesRepository) DeleteAllFavoritesByUserId(userId int) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM favorites WHERE user_id = $1", userId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (db *FavoritesRepository) DeleteFavoritesById(favoriteId int) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM favorites WHERE id = $1", favoriteId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return nil
	}
	return nil

}
