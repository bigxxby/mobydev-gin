package age

import (
	"fmt"
)

func (db *AgeRepository) GetAgeCategoryById(id int) (*AgeCategory, error) {
	query := "SELECT id, user_id, name, note, min_age, max_age FROM age_categories WHERE id = $1"
	var ageCategory AgeCategory
	err := db.Database.QueryRow(query, id).Scan(&ageCategory.ID, &ageCategory.UserID, &ageCategory.Name, &ageCategory.Note, &ageCategory.MinAge, &ageCategory.MaxAge)
	if err != nil {
		return nil, err
	}
	return &ageCategory, nil

}
func (db *AgeRepository) GetAllAgeCategories() ([]AgeCategory, error) {
	query := "SELECT id, user_id, name, note, min_age, max_age FROM age_categories"
	rows, err := db.Database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ageCategories []AgeCategory
	for rows.Next() {
		var ageCategory AgeCategory
		if err := rows.Scan(&ageCategory.ID, &ageCategory.UserID, &ageCategory.Name, &ageCategory.Note, &ageCategory.MinAge, &ageCategory.MaxAge); err != nil {
			return nil, err
		}
		ageCategories = append(ageCategories, ageCategory)
	}
	return ageCategories, nil
}
func (db *AgeRepository) GetAgeCategoryByName(name string) (*AgeCategory, error) {
	query := "SELECT id, user_id, name, note, min_age, max_age FROM age_categories WHERE name = $1"
	var ageCategory AgeCategory
	err := db.Database.QueryRow(query, name).Scan(&ageCategory.ID, &ageCategory.UserID, &ageCategory.Name, &ageCategory.Note, &ageCategory.MinAge, &ageCategory.MaxAge)
	if err != nil {
		return nil, err
	}
	return &ageCategory, nil
}

func (db *AgeRepository) getAgeCategoryByAge(query string, minAge, maxAge int) ([]*AgeCategory, error) {
	rows, err := db.Database.Query(query, minAge, maxAge)
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	var ageCategories []*AgeCategory
	for rows.Next() {
		var ageCategory AgeCategory
		if err := rows.Scan(&ageCategory.ID, &ageCategory.UserID, &ageCategory.Name, &ageCategory.Note, &ageCategory.MinAge, &ageCategory.MaxAge); err != nil {
			return nil, fmt.Errorf("failed to scan age category: %v", err)
		}
		ageCategories = append(ageCategories, &ageCategory)
	}

	return ageCategories, nil
}

func (db *AgeRepository) GetAgeCategoryByMinAge(minAge int) ([]*AgeCategory, error) {
	query := "SELECT id, user_id, name, note, min_age, max_age FROM age_categories WHERE min_age >= $1"
	return db.getAgeCategoryByAge(query, minAge, 999)
}

func (db *AgeRepository) GetAgeCategoryByMaxAge(maxAge int) ([]*AgeCategory, error) {
	query := "SELECT id, user_id, name, note, min_age, max_age FROM age_categories WHERE max_age <= $1"
	return db.getAgeCategoryByAge(query, 0, maxAge)
}

func (db *AgeRepository) GetAgeCategoryByAge(minAge, maxAge int) ([]*AgeCategory, error) {
	query := "SELECT id, user_id, name, note, min_age, max_age FROM age_categories WHERE min_age >= $1 AND max_age <= $2"
	return db.getAgeCategoryByAge(query, minAge, maxAge)
}
