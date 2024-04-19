package datasets

import (
	"project/internal/database"
)

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
			role TEXT NOT NULL,
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
		year INTEGER NOT NULL,
		category_id INTEGER NOT NULL REFERENCES categories(id),
		age_category_id INTEGER NOT NULL REFERENCES age_categories(id),
		genre_id INTEGER NOT NULL REFERENCES genres(id),
		duration_minutes INTEGER NOT NULL,
		keywords TEXT NOT NULL,
		description TEXT NOT NULL,
		director TEXT NOT NULL,
		producer TEXT NOT NULL, 
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		
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

func CreateSeasonsTable(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS seasons (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL REFERENCES users(id),
		movie_id INTEGER REFERENCES movies(id),  
		season_number INTEGER NOT NULL,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		release_date DATE NOT NULL
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
		user_id INTEGER NOT NULL REFERENCES users(id),
		url TEXT NOT NULL, 
		season_id INTEGER REFERENCES seasons(id),  
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
func CreateTrendsTable(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS trends (
		id SERIAL PRIMARY KEY,
		movie_id INTEGER REFERENCES movies(id),
		trend_date DATE NOT NULL,
		trend_value INTEGER NOT NULL
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
func CreateFavoritesTable(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS favorites (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL REFERENCES users(id),
		movie_id INTEGER NOT NULL REFERENCES movies(id),
		added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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
func CreateGenresTable(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS genres (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL REFERENCES users(id), 
		name TEXT NOT NULL UNIQUE,
		description TEXT NOT NULL
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

func CreateCategoriesTable(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL REFERENCES users(id), 
		name TEXT NOT NULL UNIQUE,
		description TEXT NOT NULL
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

func CreateAgeCategoriesTable(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS age_categories (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL REFERENCES users(id), 
		name TEXT NOT NULL UNIQUE,
		note TEXT, 
		min_age INTEGER NOT NULL,
		max_age  INTEGER NOT NULL 
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
func CreateCodesTable(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS codes (
		id SERIAL PRIMARY KEY,          
		user_email TEXT NOT NULL, 
		code INT NOT NULL,
		token TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
		expires_at TIMESTAMP

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
