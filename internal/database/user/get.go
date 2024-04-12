package user

import "log"

func (db *UserRepository) GetUserById(id int) (*User, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var user User
	err := db.Database.QueryRow("SELECT id , email , name , phone , date_of_birth , role FROM users WHERE id = $1", id).Scan(&user.Id, &user.Email, &user.Name, &user.Phone, &user.DateOfBirth, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
