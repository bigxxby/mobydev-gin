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

	// err = insertTestImages(db)
	// if err != nil {
	// 	return err
	// }

	err = insertTestMovies(db)
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
        VALUES ('testuser1@example.com', 'testpassword1', 'Test User 1', '1234567890', '1990-01-01', 0),
               ('testuser2@example.com', 'testpassword2', 'Test User 2', '0987654321', '1995-02-15', 0),
               ('admin@example.com', 'adminpassword', 'Admin', '9876543210', '1985-05-20', 1)
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

// func insertTestImages(db *database.Database) error {
// 	tx, err := db.Database.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()

// 	_, err = tx.Exec(`
//         INSERT INTO images (name, url)
//         VALUES ('image1', 'https://example.com/image1.jpg'),
//                ('image2', 'https://example.com/image2.jpg')
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

func insertTestMovies(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
        INSERT INTO movies (user_id, image_url, name, category, project_type, year, age_category, duration_minutes, keywords, description, director, producer)
        VALUES (1, 'https://www.prolydian.com/sites/default/files/2020-12/api.png', 'Movie 1', 'Action', 'Feature Film', 2020, 'PG-13', 120, 'action, thriller', 'Description for Movie 1', 'Director 1', 'Producer 1'),
               (2, 'https://www.prolydian.com/sites/default/files/2020-12/api.png', 'Movie 2', 'Comedy', 'Short Film', 2019, 'PG', 90, 'comedy, romance', 'Description for Movie 2', 'Director 2', 'Producer 2')
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
        INSERT INTO seasons (movie_id, season_number, name, description, release_date)
        VALUES (1, 1, 'Season 1', 'Description for Season 1', '2021-01-01'),
               (1, 2, 'Season 2', 'Description for Season 2', '2022-01-01')
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
        VALUES (1, 1, 'Episode 1', 30, '2021-01-10', 'Description for Episode 1'),
               (1, 2, 'Episode 2', 25, '2021-01-17', 'Description for Episode 2')
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

// func GenerateTestData(db *database.Database) error {
// 	// Генерация пользователей
// 	for i := 0; i < 10; i++ {
// 		email := fmt.Sprintf("user%d@example.com", i+1)
// 		password := "password123"
// 		name := fmt.Sprintf("User %d", i+1)
// 		phone := fmt.Sprintf("123-456-%d", i+1)
// 		dateOfBirth := time.Now().AddDate(-rand.Intn(50)+1970, rand.Intn(12)+1, rand.Intn(28)+1).Format("2006-01-02")
// 		isAdmin := rand.Intn(2) // 0 или 1

// 		_, err := db.Database.Exec(`
// 			INSERT INTO users (email, password, name, phone, date_of_birth, is_admin)
// 			VALUES ($1, $2, $3, $4, $5, $6)
// 		`, email, password, name, phone, dateOfBirth, isAdmin)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	// Генерация изображений
// 	for i := 0; i < 10; i++ {
// 		name := fmt.Sprintf("image%d.jpg", i+1)
// 		url := "https://www.prolydian.com/sites/default/files/2020-12/api.png"

// 		_, err := db.Database.Exec(`
// 			INSERT INTO images (name, url)
// 			VALUES ($1, $2)
// 		`, name, url)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	// Генерация фильмов
// 	for i := 0; i < 10; i++ {
// 		name := fmt.Sprintf("Movie %d", i+1)
// 		userID := rand.Intn(10) + 1
// 		imageID := rand.Intn(10) + 1
// 		category := randCategory()
// 		projectType := randProjectType()
// 		year := rand.Intn(24) + 2000
// 		ageCategory := randAgeCategory()
// 		durationMinutes := rand.Intn(120) + 60
// 		keywords := fmt.Sprintf("keyword%d", i+1)
// 		description := fmt.Sprintf("Description for Movie %d", i+1)
// 		director := fmt.Sprintf("Director %d", i+1)
// 		producer := fmt.Sprintf("Producer %d", i+1)

// 		_, err := db.Database.Exec(`
// 			INSERT INTO movies (name, user_id, image_id, category, project_type, year, age_category, duration_minutes, keywords, description, director, producer)
// 			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
// 		`, name, userID, imageID, category, projectType, year, ageCategory, durationMinutes, keywords, description, director, producer)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	// Генерация сезонов и эпизодов (для одного фильма, чтобы не усложнять пример)
// 	for seasonNumber := 1; seasonNumber <= 3; seasonNumber++ {
// 		seasonName := fmt.Sprintf("Season %d", seasonNumber)
// 		seasonDescription := fmt.Sprintf("Description for Season %d", seasonNumber)
// 		releaseDate := time.Now().AddDate(0, -rand.Intn(12), rand.Intn(28)).Format("2006-01-02")
// 		var seasonID int

// 		err := db.Database.QueryRow(`
// 			INSERT INTO seasons (movie_id, season_number, name, description, release_date)
// 			VALUES (1, $1, $2, $3, $4) RETURNING id
// 		`, seasonNumber, seasonName, seasonDescription, releaseDate).Scan(&seasonID)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		for episodeNumber := 1; episodeNumber <= 5; episodeNumber++ {
// 			episodeName := fmt.Sprintf("Episode %d", episodeNumber)
// 			durationMinutes := rand.Intn(40) + 20
// 			releaseDate := time.Now().AddDate(0, -rand.Intn(12), rand.Intn(28)).Format("2006-01-02")
// 			episodeDescription := fmt.Sprintf("Description for Episode %d", episodeNumber)

// 			_, err := db.Database.Exec(`
// 				INSERT INTO episodes (season_id, episode_number, name, duration_minutes, release_date, description)
// 				VALUES ($1, $2, $3, $4, $5, $6)
// 			`, seasonID, episodeNumber, episodeName, durationMinutes, releaseDate, episodeDescription)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 		}
// 	}

//		return nil
//	}
func DropTables(db *sql.DB) error {
	tables := []string{"episodes", "seasons", "movies", "images", "users"}

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
