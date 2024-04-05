package init

import (
	"database/sql"
	"log"
	"project/internal/database"

	_ "github.com/lib/pq"
)

func CreateConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=db port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")

	if err != nil {
		return nil, err
	}

	log.Println("trying to connect...")
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database connected")
	return db, nil
}

func CreateDatabaseStruct() (*database.Database, error) {
	db, err := CreateConnection()
	if err != nil {
		return nil, err
	}
	manager := database.Database{
		Database: db,
	}
	return &manager, err
}
