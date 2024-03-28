package database

import (
	"log"
	"project/internal/logic"

	"golang.org/x/crypto/bcrypt"
)

func (db *Database) CheckUser(email, password string) (*User, bool, error) {
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

		sessionID, err := logic.GenerateSessionID()
		if err != nil {
			return nil, false, err
		}

		user.SessionId.String = sessionID

		_, err = tx.Exec("UPDATE users SET session_id = $1 WHERE id = $2", sessionID, user.Id)
		if err != nil {
			return nil, false, err
		}

		err = tx.Commit()
		if err != nil {
			log.Println(err.Error())
			return nil, false, err
		}

		return user, true, nil
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
		return nil, false, err
	}
	return nil, false, nil
}

func (db *Database) IsAuthorised(sessionId string) (*User, bool, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	tx, err := db.Database.Begin()
	if err != nil {
		log.Println(err.Error())
		return nil, false, err
	}

	stmt, err := tx.Prepare("SELECT * FROM users WHERE session_id = $1")
	if err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return nil, false, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(sessionId)
	if err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return nil, false, err
	}
	defer rows.Close()

	user := &User{}

	if rows.Next() {
		if err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Phone, &user.DateOfBirth, &user.SessionId, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
			tx.Rollback()
			log.Println(err.Error())
			return nil, false, err
		}
		if err := tx.Commit(); err != nil {
			log.Println(err.Error())
			return nil, false, err
		}
		return user, true, nil
	}

	if err := tx.Commit(); err != nil {
		log.Println(err.Error())
		return nil, false, err
	}

	return nil, false, nil
}

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
