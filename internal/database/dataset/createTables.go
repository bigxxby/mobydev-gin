package database

import (
	"log"
	"project/internal/database"
)

func CreateTables(db *database.Database) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := CreateUsersTable(db)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = CreateGenresTable(db)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = CreateCategoriesTable(db)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = CreateAgeCategoriesTable(db)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = CreateMoviesTable(db)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = CreateSeasonsTable(db)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = CreateEpisodesTable(db)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = CreateTrendsTable(db)
	if err != nil {
		return err
	}
	err = CreateFavoritesTable(db)
	if err != nil {
		return err
	}
	return nil
}
