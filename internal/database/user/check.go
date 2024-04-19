package user

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (db *UserRepository) CheckUserExistsByEmail(email string) error {
	stmt, err := db.Database.Prepare("SELECT * FROM users WHERE email = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(email)
	if err != nil {
		return err
	}
	defer rows.Close()

	exists := rows.Next()
	if !exists {
		return sql.ErrNoRows
	}
	return nil
}
func (db *UserRepository) CheckUserExistsById(id int) error {
	stmt, err := db.Database.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return err
	}
	defer rows.Close()

	exists := rows.Next()
	if !exists {
		return sql.ErrNoRows
	}
	return nil
}
func (db *UserRepository) CheckUserСredentials(email, password string) (*User, bool, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	stmt, err := db.Database.Prepare("SELECT * FROM users WHERE email = $1")
	if err != nil {
		return nil, false, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(email)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()

	user := &User{}

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Phone, &user.DateOfBirth, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, false, err
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			return nil, false, nil
		}

		return user, true, nil
	}
	return nil, false, nil
}
