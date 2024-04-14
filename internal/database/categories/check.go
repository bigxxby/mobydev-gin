package categories

import (
	"log"
)

func (db *CategoryRepository) CheckCategoryExistsById(categoryId int) (bool, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	stmt, err := db.Database.Prepare("SELECT * FROM categories WHERE id = $1")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(categoryId)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	exists := rows.Next()

	return exists, nil
}
func (db *CategoryRepository) CheckCategoryExistsByName(name string) (bool, error) {
	var exists bool

	query := `
        SELECT EXISTS(
            SELECT 1 FROM categories WHERE name = $1
        )
    `

	err := db.Database.QueryRow(query, name).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
func (db *CategoryRepository) CheckCategoryIsUsedInMovies(categoryId int) (bool, error) {
	var exists bool

	query := `
        SELECT EXISTS(
            SELECT 1 FROM movies WHERE category_id = $1
        )
    `

	err := db.Database.QueryRow(query, categoryId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
