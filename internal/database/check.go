package database

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// checks if user exists and returns it if true
func (db *Database) CheckUserExists(email, password string) (*User, bool, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	tx, err := db.Database.Begin()
	if err != nil {
		return nil, false, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("SELECT * FROM users WHERE email = $1")
	if err != nil {
		return nil, false, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(email)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()

	user := &User{}

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Phone, &user.DateOfBirth, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, false, err
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			return nil, false, nil
		}

		return user, true, nil
	}
	err = tx.Commit()
	if err != nil {
		return nil, false, err
	}
	return nil, false, nil
}
