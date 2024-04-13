package utils

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func CreateJWTToken(userId int, email, name string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId,
			"email":  email,
			"name":   name,
			"role":   role,
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func VerifyToken(tokenString string) (int, string, error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims) //get user id
	if !ok {
		return 0, "", errors.New("invalid token claims")
	}

	userId, ok := claims["userId"].(float64)
	if !ok {
		return 0, "", errors.New("invalid userId")
	}
	userRole, ok := claims["role"].(string)
	if !ok {
		return 0, "", errors.New("invalid role")
	}

	return int(userId), userRole, nil
}
