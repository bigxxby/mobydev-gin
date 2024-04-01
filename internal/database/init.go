package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func CreateConnection() (*sql.DB, error) {
	// var connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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

func CreateDatabaseStruct() (*Database, error) {
	db, err := CreateConnection()
	if err != nil {
		return nil, err
	}
	manager := Database{
		Database: db,
	}
	return &manager, err
}
