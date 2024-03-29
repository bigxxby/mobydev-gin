package database

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// checks if user exists and returns it if true
func (db *Database) CheckIfUserExists(email, password string) (*User, bool, error) {
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
		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Phone, &user.DateOfBirth, &user.SessionId, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
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

func (d *Database) AuthoreseUserById(id int, sessionId string) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE users SET session_id = $1 WHERE id = $2", sessionId, id)

	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// func (db *Database) FindUserBySessionId(sessionId string) (*User, error) { // CONITNUEEEEEEEEEEEEEEEEE
// }

// func (db *Database) IsAuthorised(sessionId string) (*User, bool, error) {
// 	log.SetFlags(log.LstdFlags | log.Lshortfile)
// 	tx, err := db.Database.Begin()
// 	if err != nil {
// 		log.Println(err.Error())
// 		return nil, false, err
// 	}
// 	defer tx.Rollback()
// 	stmt, err := tx.Prepare("SELECT * FROM users WHERE session_id = $1")
// 	if err != nil {
// 		log.Println(err.Error())
// 		return nil, false, err
// 	}
// 	defer stmt.Close()

// 	rows, err := stmt.Query(sessionId)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return nil, false, err
// 	}
// 	defer rows.Close()

// 	user := &User{}

// 	if rows.Next() {
// 		if err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Phone, &user.DateOfBirth, &user.SessionId, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
// 			tx.Rollback()
// 			log.Println(err.Error())
// 			return nil, false, err
// 		}
// 		if err := tx.Commit(); err != nil {
// 			log.Println(err.Error())
// 			return nil, false, err
// 		}
// 		return user, true, nil
// 	}

// 	if err := tx.Commit(); err != nil {
// 		log.Println(err.Error())
// 		return nil, false, err
// 	}

// 	return nil, false, nil
// }

func (db *Database) LogoutUser(sessionId string) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	tx, err := db.Database.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			log.Println("Transaction rolled back due to error:", err.Error())
		}
	}()

	_, err = tx.Exec("UPDATE users SET session_id = '' WHERE session_id = $1", sessionId)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Error committing transaction:", err.Error())
		return err
	}

	return nil
}
