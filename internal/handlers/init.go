package handlers

import "project/internal/database"

type Manager struct {
	DB *database.Database
}

func Init() (*Manager, error) {
	manager := Manager{}
	var err error
	manager.DB, err = database.CreateDatabaseStruct()
	if err != nil {
		return nil, err
	}
	return &manager, nil
}
