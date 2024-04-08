package utils

import (
	"errors"
	"net/mail"
	"regexp"
	"strconv"
	"unicode"
)

func CheckValidForReg(email, password, confirmPassword string) error {
	if password != confirmPassword {
		return errors.New("passwords Does Not Match")
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}
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
