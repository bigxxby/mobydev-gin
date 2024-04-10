package database

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// checks if user with this Сredentials exists
func (db *Database) CheckUserСredentials(email, password string) (*User, bool, error) {
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
		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Phone, &user.DateOfBirth, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
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

// checks if user really exists
func (db *Database) CheckUserExistsById(Id string) (bool, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	tx, err := db.Database.Begin()
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(Id)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	exists := rows.Next()

	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (db *Database) CheckProjectExistsById(projectId int) (bool, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	tx, err := db.Database.Begin()
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("SELECT * FROM projects WHERE id = $1")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(projectId)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	exists := rows.Next()

	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return exists, nil
}
