package database

import "log"

func (db *Database) GetUserById(id int) (*User, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	tx, err := db.Database.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	var user User
	err = tx.QueryRow("SELECT id , email , name , phone , date_of_birth , is_admin FROM users WHERE id = $1", id).Scan(&user.Id, &user.Email, &user.Name, &user.Phone, &user.DateOfBirth, &user.IsAdmin)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (db *Database) GetMovies(limit int) ([]Movie, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if limit != 0 {
		tx, err := db.Database.Begin()
		if err != nil {
			return nil, err
		}
		defer tx.Rollback()

		stmt, err := tx.Prepare("SELECT * FROM movies LIMIT $1")
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		rows, err := stmt.Query(limit)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var movies []Movie

		for rows.Next() {
			var movie Movie
			err := rows.Scan(
				&movie.Id,
				&movie.UserId,
				&movie.ImageUrl,
				&movie.Name,
				&movie.Category,
				&movie.ProjectType,
				&movie.Year,
				&movie.AgeCategory,
				&movie.DurationMinutes,
				&movie.Keywords,
				&movie.Description,
				&movie.Director,
				&movie.Producer,
			)
			if err != nil {
				return nil, err
			}
			movies = append(movies, movie)
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}

		err = tx.Commit()
		if err != nil {
			return nil, err
		}

		return movies, nil

	} else {

		tx, err := db.Database.Begin()
		if err != nil {
			return nil, err
		}
		defer tx.Rollback()

		stmt, err := tx.Prepare("SELECT * FROM movies ")
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		rows, err := stmt.Query()
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var movies []Movie

		for rows.Next() {
			var movie Movie
			err := rows.Scan(
				&movie.Id,
				&movie.UserId,
				&movie.ImageUrl,
				&movie.Name,
				&movie.Category,
				&movie.ProjectType,
				&movie.Year,
				&movie.AgeCategory,
				&movie.DurationMinutes,
				&movie.Keywords,
				&movie.Description,
				&movie.Director,
				&movie.Producer,
			)
			if err != nil {
				return nil, err
			}
			movies = append(movies, movie)
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}

		err = tx.Commit()
		if err != nil {
			return nil, err
		}

		return movies, nil
	}
}
