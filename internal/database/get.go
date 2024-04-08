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
func (db *Database) GetProjects(limit int) ([]Project, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if limit != 0 {
		tx, err := db.Database.Begin()
		if err != nil {
			return nil, err
		}
		defer tx.Rollback()

		stmt, err := tx.Prepare("SELECT * FROM projects  ORDER BY created_at LIMIT $1")
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		rows, err := stmt.Query(limit)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var projects []Project

		for rows.Next() {
			var project Project
			err := rows.Scan(
				&project.Id,
				&project.UserId,
				&project.ImageUrl,
				&project.Name,
				&project.Category,
				&project.ProjectType,
				&project.Year,
				&project.AgeCategory,
				&project.DurationMinutes,
				&project.Keywords,
				&project.Description,
				&project.Director,
				&project.Producer,
				&project.CreatedAt,
				&project.UpdatedAt,
			)
			if err != nil {
				return nil, err
			}
			projects = append(projects, project)
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}

		err = tx.Commit()
		if err != nil {
			return nil, err
		}

		return projects, nil

	} else {

		tx, err := db.Database.Begin()
		if err != nil {
			return nil, err
		}
		defer tx.Rollback()

		stmt, err := tx.Prepare("SELECT * FROM projects ORDER BY created_at")
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		rows, err := stmt.Query()
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var projects []Project

		for rows.Next() {
			var project Project
			err := rows.Scan(
				&project.Id,
				&project.UserId,
				&project.ImageUrl,
				&project.Name,
				&project.Category,
				&project.ProjectType,
				&project.Year,
				&project.AgeCategory,
				&project.DurationMinutes,
				&project.Keywords,
				&project.Description,
				&project.Director,
				&project.Producer,
				&project.CreatedAt,
				&project.UpdatedAt,
			)
			if err != nil {
				return nil, err
			}
			projects = append(projects, project)
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}

		err = tx.Commit()
		if err != nil {
			return nil, err
		}

		return projects, nil
	}
}
func (db *Database) GetProjectById(id int) (*Project, error) {

	tx, err := db.Database.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	var project Project
	err = tx.QueryRow("SELECT * FROM projects WHERE id = $1", id).Scan(
		&project.Id,
		&project.UserId,
		&project.ImageUrl,
		&project.Name,
		&project.Category,
		&project.ProjectType,
		&project.Year,
		&project.AgeCategory,
		&project.DurationMinutes,
		&project.Keywords,
		&project.Description,
		&project.Director,
		&project.Producer,
		&project.CreatedAt,
		&project.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &project, nil
}
