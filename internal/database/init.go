package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func CreateConnection() (*sql.DB, error) {
	connStr := "user=postgres password=postgres host=test-database.c12yiigcoe0w.eu-north-1.rds.amazonaws.com port=5432" // use env next time
	db, err := sql.Open("postgres", connStr)
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
