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
	return nil
}
func CreateUsersTable(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            email TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL,
			name TEXT, 
            phone TEXT,
            date_of_birth DATE,
			is_admin INTEGER,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            deleted_at TIMESTAMP
        )
    `)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func CreateMoviesTable(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS movies (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL REFERENCES users(id),
		image_url TEXT NOT NULL,
		name TEXT NOT NULL,
		category TEXT NOT NULL,
		project_type TEXT NOT NULL,
		year INTEGER NOT NULL,
		age_category TEXT NOT NULL,
		duration_minutes INTEGER NOT NULL,
		keywords TEXT NOT NULL,
		description TEXT NOT NULL,
		director TEXT NOT NULL,
		producer TEXT NOT NULL
	);
	
    `)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// func CreateImagesTable(db *database.Database) error {
// 	tx, err := db.Database.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()

// 	_, err = tx.Exec(`
//         CREATE TABLE IF NOT EXISTS images (
// 			id SERIAL PRIMARY KEY,
// 			name VARCHAR(255),
// 			url VARCHAR(255)
// 		)
//     `)
// 	if err != nil {
// 		return err
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		return err
// 	}

//		return nil
//	}
func CreateSeasonsTable(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS seasons (
		id SERIAL PRIMARY KEY,
		movie_id INTEGER REFERENCES movies(id),  
		season_number INTEGER NOT NULL,
		name TEXT,
		description TEXT,
		release_date DATE
	);
    `)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
func CreateEpisodesTable(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS episodes (
		id SERIAL PRIMARY KEY,
		season_id INTEGER REFERENCES seasons(id),  -- Внешний ключ на таблицу seasons
		episode_number INTEGER NOT NULL,
		name TEXT,
		duration_minutes INTEGER,
		release_date DATE,
		description TEXT
	);
    `)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// func CreateRolesTable(db *database.Database) error {
// 	tx, err := db.Database.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()

// 	_, err = tx.Exec(`
//         CREATE TABLE IF NOT EXISTS roles (
// 			id SERIAL PRIMARY KEY,
// 			role TEXT

// 		)
//     `)
// 	if err != nil {
// 		return err
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func CreateAdmin(db *database.Database) error {
// 	tx, err := db.Database.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()

// 	password := "12345678Aa#"

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = tx.Exec(`
//         INSERT INTO users (email, password, name, is_admin, created_at, updated_at)
//         VALUES ($1, $2, $3, $4, NOW(), NOW())
//     `, "admin@example.com", hashedPassword, "Admin", 1)
// 	if err != nil {
// 		return err
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
