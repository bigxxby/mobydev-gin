package datasets

import (
	"project/internal/database"
)

func InitDatasets(database *database.Database) error {

	//test data
	err := DropTables(database.Database)
	if err != nil {
		return err
	}
	err = CreateTables(database)
	if err != nil {
		return err
	}
	err = CreateTestData(database)
	if err != nil {
		return err
	}
	return nil
}
