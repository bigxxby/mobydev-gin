package database

import (
	"errors"
	"log"
	"project/internal/utils"
)

func (db *Database) CreateUser(email, password string, role string) (bool, error) {
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
func (db *Database) CreateProject(userId int, imageUrl string, name string, category string, projectType string, year int, AgeCategory string, durationMinutes int, keywords string, desc string, director string, producer string) (*Project, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Start a new transaction
	tx, err := db.Database.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Prepare the SQL statement for inserting a new project
	stmt, err := tx.Prepare(`
		INSERT INTO projects (
			user_id, image_url, name, category, project_type, year, age_category, duration_minutes, keywords, description, director, producer
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id, created_at, updated_at
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement
	var project Project
	err = stmt.QueryRow(
		userId, imageUrl, name, category, projectType, year, AgeCategory, durationMinutes, keywords, desc, director, producer,
	).Scan(
		&project.Id, &project.CreatedAt, &project.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &project, nil
}
