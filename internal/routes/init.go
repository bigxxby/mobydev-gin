package routes

import (
	"project/internal/database"
	i "project/internal/init"
)

type Manager struct {
	DB *database.Database
}

func Init() (*Manager, error) {
	manager := Manager{}
	var err error
	manager.DB, err = i.CreateDatabaseStruct()
	if err != nil {
		return nil, err
	}
	return &manager, nil
}
