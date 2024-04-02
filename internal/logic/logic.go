package logic

import (
	"errors"
	"net/mail"
	"regexp"
	"unicode"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	onlyLatin := containsNonLatinLetters(password)
	hasUpperCase := regexp.MustCompile(`[A-Z]`).MatchString
	hasLowerCase := regexp.MustCompile(`[a-z]`).MatchString
	hasDigit := regexp.MustCompile(`\d`).MatchString
	hasSpecialChar := regexp.MustCompile(`[@#$%^&+=!]`).MatchString

	return !onlyLatin && hasUpperCase(password) && hasLowerCase(password) && hasDigit(password) && hasSpecialChar(password)
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

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GenerateSessionID() (string, error) {
	sessionID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return sessionID.String(), nil
}

func CheckValidForReg(email, password, confirmPassword string) error {
	if password != confirmPassword {
		return errors.New("Passwords Does Not Match")
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}
	if !IsValidPassword(password) {
		return errors.New("Password is not valid")
	}
	return nil
}
