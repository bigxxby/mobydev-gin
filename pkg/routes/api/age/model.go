package age

import (
	"project/internal/database"
	"project/internal/database/age"
)

type AgeRoute struct {
	DB *database.Database
}

func isValidAgeCategory(ageCategory age.AgeCategory) bool {
	if ageCategory.MinAge <= 0 || ageCategory.MaxAge > 100 {
		return false
	}
	if ageCategory.MaxAge > 100 {
		return false
	}
	if ageCategory.MinAge > ageCategory.MaxAge {
		return false
	}
	if ageCategory.Name == "" {
		return false
	}

	return true
}
