package utils

import (
	"errors"
	"math/rand"
	"net/mail"
	"project/internal/database/age"
	"regexp"
	"strconv"
	"time"
	"unicode"
)

func CheckValidForReg(email, password string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}
	if !isValidPassword(password) {
		return errors.New("password is not valid")
	}

	return nil
}
func CheckPasswordIsValid(password string) error {
	if !isValidPassword(password) {
		return errors.New("password is not valid")
	}
	return nil
}

func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	if len(password) > 16 {
		return false
	}
	onlyLatin := containsNonLatinLetters(password)
	hasUpperCase := regexp.MustCompile(`[A-Z]`).MatchString
	hasLowerCase := regexp.MustCompile(`[a-z]`).MatchString
	hasDigit := regexp.MustCompile(`\d`).MatchString
	hasSpecialChar := regexp.MustCompile(`[@#$%^&+=!]`).MatchString

	return !onlyLatin && hasUpperCase(password) && hasLowerCase(password) && hasDigit(password) && hasSpecialChar(password)
}

func GenerateTempPassword(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var passwordChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%^&+=")
	var specialChars = []rune("#$%&")

	// At least one uppercase letter and one special character
	hasUpper := false
	hasSpecial := false

	password := make([]rune, length)
	for i := range password {
		if i == length-2 && !hasUpper {
			password[i] = rune('A' + r.Intn(26))
			hasUpper = true
		} else if i == length-1 && !hasSpecial {
			password[i] = specialChars[r.Intn(len(specialChars))]
			hasSpecial = true
		} else {
			password[i] = passwordChars[r.Intn(len(passwordChars))]
		}
	}
	r.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})
	return string(password)
}

func containsNonLatinLetters(str string) bool { // checks if password has not latin symbols
	for _, char := range str {
		if !unicode.IsLetter(char) {
			continue
		}
		if !unicode.Is(unicode.Latin, char) {
			return true
		}
	}
	return false
}

func IsValidNum(id string) (bool, int) {
	num, err := strconv.Atoi(id)
	if err != nil {
		return false, 0
	}
	if num < 0 {
		return false, 0
	}
	if !(strconv.Itoa(num) == id) {
		return false, 0
	}

	return true, num
}

func IsValidPhoneNumber(phoneNumber string) bool {
	pattern := `^\+[1-9]\d{1,14}$`

	match, _ := regexp.MatchString(pattern, phoneNumber)
	return match
}
func IsValidAgeCategory(ageCategory age.AgeCategory) bool {
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
