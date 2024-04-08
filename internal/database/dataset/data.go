package database

import (
	"database/sql"
	"log"
	"project/internal/database"
)

func CreateTestData(db *database.Database) error {
	err := insertTestUsers(db)
	if err != nil {
		return err
	}

	err = insertTestProjects(db)
	if err != nil {
		return err
	}

	err = insertTestSeasons(db)
	if err != nil {
		return err
	}

	err = insertTestEpisodes(db)
	if err != nil {
		return err
	}

	return nil
}

func insertTestUsers(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
        INSERT INTO users (email, password, name, phone, date_of_birth, is_admin)
        VALUES 
        ('testuser1@example.com', 'testpassword1', 'Test User 1', '1234567890', '1990-01-01', 0),
        ('testuser2@example.com', 'testpassword2', 'Test User 2', '0987654321', '1995-02-15', 0),
        ('admin@example.com', 'adminpassword', 'Admin', '9876543210', '1985-05-20', 1),
		('admin2@example.com', 'adminpassword', 'Admin', '9876543210', '1985-05-20', 1)
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
func insertTestProjects(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
        INSERT INTO projects (user_id, image_url, name, category, project_type, year, age_category, duration_minutes, keywords, description, director, producer)
        VALUES 
        (1, 'https://www.prolydian.com/sites/default/files/2020-12/api.png', 'Movie 1', 'Action', 'Feature Film', 2020, 'PG-13', 120, 'action, thriller', 'Description for Movie 1', 'Director 1', 'Producer 1'),
        (2, 'https://www.prolydian.com/sites/default/files/2020-12/api.png', 'Movie 2', 'Comedy', 'Short Film', 2019, 'PG', 90, 'comedy, romance', 'Description for Movie 2', 'Director 2', 'Producer 2'),
        (3, 'https://www.prolydian.com/sites/default/files/2020-12/api.png', 'Movie 3', 'Drama', 'Feature Film', 2021, 'R', 110, 'drama', 'Description for Movie 3', 'Director 3', 'Producer 3'),
        (4, 'https://www.prolydian.com/sites/default/files/2020-12/api.png', 'Movie 4', 'Thriller', 'Feature Film', 2018, 'R', 105, 'thriller, mystery', 'Description for Movie 4', 'Director 4', 'Producer 4')
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

func insertTestSeasons(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
        INSERT INTO seasons (project_id, season_number, name, description, release_date)
        VALUES 
        (1, 1, 'Season 1', 'Description for Season 1', '2021-01-01'),
        (1, 2, 'Season 2', 'Description for Season 2', '2022-01-01'),
        (2, 1, 'Season 1', 'Description for Season 1 of Movie 2', '2020-05-01'),
        (3, 1, 'Season 1', 'Description for Season 1 of Movie 3', '2021-07-01')
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

func insertTestEpisodes(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
        INSERT INTO episodes (season_id, episode_number, name, duration_minutes, release_date, description)
        VALUES 
        (1, 1, 'Episode 1', 30, '2021-01-10', 'Description for Episode 1'),
        (1, 2, 'Episode 2', 25, '2021-01-17', 'Description for Episode 2'),
        (2, 1, 'Episode 1 of Season 2', 35, '2022-01-15', 'Description for Episode 1 of Season 2'),
        (3, 1, 'Episode 1 of Movie 3', 40, '2021-07-10', 'Description for Episode 1 of Movie 3')
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

func DropTables(db *sql.DB) error {
	tables := []string{"episodes", "seasons", "projects", "users"}

	for _, table := range tables {
		_, err := db.Exec("DROP TABLE IF EXISTS " + table + " CASCADE")
		if err != nil {
			log.Fatalf("Failed to drop table %s: %v", table, err)
			return err
		}
	}

	log.Println("All tables dropped successfully")
	return nil
}
