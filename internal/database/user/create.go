package user

import (
	"database/sql"
	"errors"
	"log"
	"project/internal/utils"
	"time"
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

func (db *UserRepository) AddVerificationCode(email, code string) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	expiresAt := time.Now().Add(time.Hour)

	q := "INSERT INTO codes (user_email, code, expires_at) VALUES ($1, $2, $3)"
	_, err = tx.Exec(q, email, code, expiresAt)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (db *UserRepository) ValidateCode(email, code string) error {
	q := `
	SELECT code
	FROM codes 
	WHERE user_email = $1 AND code = $2 AND expires_at > CURRENT_TIMESTAMP
	ORDER BY created_at DESC
	LIMIT 1;
	`
	var retrievedCode string
	err := db.Database.QueryRow(q, email, code).Scan(&retrievedCode)
	if err != nil {
		return err
	}

	if retrievedCode != code {
		return errors.New("code does not match or is expired")
	}

	return nil
}

func (db *UserRepository) SaveResetToken(token string, userEmail string) error {

	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := `UPDATE codes 
	SET token = $1 
	WHERE user_email = $2 AND code = $3 AND expires_at > CURRENT_TIMESTAMP 
	ORDER BY created_at DESC 
	LIMIT 1;`

	_, err = tx.Exec(q, token, userEmail)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}
func (db *UserRepository) VerifyToken(token string) error {
	var count int
	q := `
		SELECT COUNT(*) 
		FROM codes 
		WHERE token = ? AND expires_at > CURRENT_TIMESTAMP
	`

	err := db.Database.QueryRow(q, token).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		return sql.ErrNoRows
	}

	return nil
}
