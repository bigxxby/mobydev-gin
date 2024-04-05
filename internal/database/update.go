package database

import (
	"log"
	"time"
)

func (db *Database) UpdateUserByAdmin(userID string, newName, newPhone, newDOB string, isAdmin bool) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	tx, err := db.Database.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer tx.Rollback()

	var isAdminValue int
	if isAdmin {
		isAdminValue = 1
	} else {
		isAdminValue = 0
	}
	now := time.Now()
	_, err = tx.Exec("UPDATE users SET name=$1, phone=$2, date_of_birth=$3, is_admin=$4, updated_at=$5 WHERE id=$6",
		newName, newPhone, newDOB, isAdminValue, now, userID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
