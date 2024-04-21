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

	query := "UPDATE users SET name = $1 ,  phone = $2 , date_of_birth = $3 ,updated_at = CURRENT_TIMESTAMP   WHERE id = $4"
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
func (db *UserRepository) UpdatePassword(email, password string) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := "UPDATE users SET password = $1, updated_at = CURRENT_TIMESTAMP  WHERE email = $2  "
	_, err = tx.Exec(query, password, email)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return nil
	}
	return nil
}
