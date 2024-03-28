package database

import (
	"log"
	"project/internal/logic"
	"time"
)

func (db *Database) AddUser(email, password string) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	hashedPassword, err := logic.HashPassword(password)
	if err != nil {
		log.Println("Error Hashing password : ", err.Error())
		return err
	}

	_, err = tx.Exec(`
	    INSERT INTO users (email, password) VALUES ($1, $2)
	`, email, string(hashedPassword))
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) UpdateUser(userID string, newName, newPhone string, newDOB string) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	now := time.Now()
	_, err = tx.Exec("UPDATE users SET name=$1, phone=$2, date_of_birth=$3, updated_at=$4 WHERE id=$5", newName, newPhone, newDOB, now, userID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) DeleteUser(userID string) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM users WHERE id=$1", userID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdateUserByAdmin(userID string, newName, newPhone, newDOB string, isAdmin bool) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	tx, err := db.Database.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer tx.Rollback()

	var isAdminValue int
	if isAdmin {
		isAdminValue = 1
	} else {
		isAdminValue = 0
	}
	now := time.Now()
	_, err = tx.Exec("UPDATE users SET name=$1, phone=$2, date_of_birth=$3, is_admin=$4, updated_at=$5 WHERE id=$6",
		newName, newPhone, newDOB, isAdminValue, now, userID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

// func CreateUserFromRegistration(email, password string) (*User, error) {
// 	hashedPassword, err := logic.HashPassword(password)
// 	if err != nil {
// 		log.Println("Error Hashing password : ", err.Error())
// 		return nil, err
// 	}
// 	user := User{
// 		Email:    email,
// 		Password: hashedPassword,
// 	}
// 	return &user, nil
// }
