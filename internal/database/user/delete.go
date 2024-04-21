package user

func (db *UserRepository) DeleteResetData(email string) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := "DELETE FROM codes WHERE user_email = $1"
	_, err = tx.Exec(q, email)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
