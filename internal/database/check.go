package database

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// checks if user with this Сredentials exists
func (db *Database) CheckUserСredentials(email, password string) (*User, bool, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	stmt, err := db.Database.Prepare("SELECT * FROM users WHERE email = $1")
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
	return nil, false, nil
}

// checks if user really exists
func (db *Database) CheckUserExistsById(id int) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	stmt, err := db.Database.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return err
	}
	defer rows.Close()

	exists := rows.Next()
	if !exists {
		return sql.ErrNoRows
	}
	return nil
}

func (db *Database) CheckMovieExistsById(movieId int) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	stmt, err := db.Database.Prepare("SELECT * FROM movies WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(movieId)
	if err != nil {
		return err
	}
	defer rows.Close()

	exists := rows.Next()
	if !exists {
		return sql.ErrNoRows
	}

	return nil
}

func (db *Database) CheckIfMovieAdded(userId, movieId int) (bool, error) {
	var exists bool

	err := db.Database.QueryRow("SELECT EXISTS(SELECT 1 FROM favorites WHERE user_id=$1 AND movie_id=$2)", userId, movieId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
