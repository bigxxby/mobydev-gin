package user

import (
	"time"
)

func (db *UserRepository) UpdateProfile(userId int, name, phone string, dob *time.Time) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var date *time.Time
	if dob != nil {
		date = dob
	}

	query := "UPDATE users SET name = $1 ,  phone = $2 , date_of_birth = $3 WHERE id = $4"
	_, err = tx.Exec(query, name, phone, date, userId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return nil
	}
	return nil
}
