package user

import (
	"errors"
	"log"
	"project/internal/utils"
)

func (db *UserRepository) CreateUser(email, password string, role string) (bool, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var count int
	err := db.Database.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&count)
	if err != nil {
		log.Println("Error checking email existence: ", err.Error())
		return false, err
	}
	if count > 0 {
		return true, errors.New("User with email already exists")
	}

	tx, err := db.Database.Begin()
	if err != nil {
		return false, err
	}
	defer tx.Rollback()
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Println("Error Hashing password : ", err.Error())
		return false, err
	}

	_, err = tx.Exec(`
	    INSERT INTO users (email, password , role ) VALUES ($1, $2 , $3)
	`, email, string(hashedPassword), role)
	if err != nil {
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return false, nil
}
