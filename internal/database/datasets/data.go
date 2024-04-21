package datasets

import (
	"project/internal/database"
	"project/internal/utils"
)

func insertTestUsers(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
        INSERT INTO users (email, password, name, phone, date_of_birth, role)
        VALUES 
        ('testuser1@example.com', 'testpassword1', 'Test User 1', '1234567890', '1990-01-01', 'admin'),
        ('testuser2@example.com', 'testpassword2', 'Test User 2', '0987654321', '1995-02-15', 'admin'),
        ('admin@example.com', 'adminpassword', 'Admin', '9876543210', '1985-05-20', 'user'),
		('admin2@example.com', 'adminpassword', 'Admin', '9876543210', '1985-05-20', 'user')
    `)
	if err != nil {
		return err
	}

	//create admin
	passAdmin, err := utils.HashPassword("Aa12345678#")
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
	INSERT INTO users (email, password, name, phone, date_of_birth, role)
	VALUES 
	('big@example.com', $1, 'Admin', '9876543210', '1985-05-20', 'admin'),
	('bigxxby@yandex.ru', $1, 'Admin', '9876543210', '1985-05-20', 'admin')`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(passAdmin)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
func insertTestMovies(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
        INSERT INTO movies (user_id, image_url,category_id,age_category_id,genre_id, name, year, duration_minutes, keywords, description, director, producer)
        VALUES 
        (1, 'https://www.prolydian.com/sites/default/files/2020-12/api.png',1,1,1, 'Movie 1',  202, 120, 'action, thriller', 'Description for Movie 1', 'Director 1', 'Producer 1'),
        (2, 'https://www.prolydian.com/sites/default/files/2020-12/api.png',1,1,1, 'Movie 2',  2019, 90, 'comedy, romance', 'Description for Movie 2', 'Director 2', 'Producer 2'),
        (3, 'https://www.prolydian.com/sites/default/files/2020-12/api.png',1,1,1, 'Movie 3',  2021,  110, 'drama', 'Description for Movie 3', 'Director 3', 'Producer 3'),
        (4, 'https://www.prolydian.com/sites/default/files/2020-12/api.png',1,1,1, 'Movie 4',  2018,  105, 'thriller, mystery', 'Description for Movie 4', 'Director 4', 'Producer 4')
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
        INSERT INTO seasons (movie_id,user_id, season_number, name, description, release_date)
        VALUES 
        (1, 1, 1, 'Season 1', 'Description for Season 1', '2021-01-01'),
        (1, 1, 2, 'Season 2', 'Description for Season 2', '2022-01-01'),
        (2, 1, 1, 'Season 1', 'Description for Season 1 of Movie 2', '2020-05-01'),
        (3, 1,1, 'Season 1', 'Description for Season 1 of Movie 3', '2021-07-01')
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
        INSERT INTO episodes (season_id,user_id, url , episode_number, name, duration_minutes, release_date, description)
        VALUES 
        (1, 1,'https://www.youtube.com/', 1, 'Episode 1', 30, '2021-01-10', 'Description for Episode 1'),
        (1, 1,'https://www.youtube.com/', 2, 'Episode 2', 25, '2021-01-17', 'Description for Episode 2'),
        (2, 1,'https://www.youtube.com/', 1, 'Episode 1 of Season 2', 35, '2022-01-15', 'Description for Episode 1 of Season 2'),
        (3, 1,'https://www.youtube.com/', 1, 'Episode 1 of Movie 3', 40, '2021-07-10', 'Description for Episode 1 of Movie 3')
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
func insertTestTrends(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
	INSERT INTO trends (movie_id, trend_date, trend_value)
	VALUES 
	(1, '2021-01-10', 1000),
	(1, '2021-01-17', 1500),
	(2, '2022-01-15', 1200),
	(3, '2021-07-10', 800)
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
func insertTestFavorites(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
	INSERT INTO favorites (user_id, movie_id)
	VALUES 
	(1, 1),
	(1, 2),
	(2, 3)
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
func insertTestCategories(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
	INSERT INTO categories (user_id, name , description)
	VALUES 
	(5, 'Series' , 'ddescription1'),
	(5, 'Movie' , 'ddescription2'),
	(5, 'Telehikaya' , 'ddescription3')
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
func insertTestAgeCategoires(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
	INSERT INTO age_categories (user_id, name, note, min_age, max_age)
	VALUES 
	(1, 'PG', 'Parental Guidance Suggested – Some material may not be suitable for children.', 13, 18),
	(1, 'PG-13', 'Parents Strongly Cautioned – Some material may be inappropriate for children under 13.', 13, 13),
	(1, 'R', 'Restricted – Under 17 requires accompanying parent or adult guardian.', 17, 17),
	(1, 'G', 'General Audiences – All ages admitted.', 0, 99);
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
func insertTestGenres(db *database.Database) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
	INSERT INTO genres (user_id ,  name , description)
	VALUES 
	(1,'Action', 'this is actiion'),
	(1,'Comedy' , 'this is comedy'),
	(1,'Drama' , 'this is drama ')
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
