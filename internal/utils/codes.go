package utils

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

func GenerateResetToken(length int) string {
	key := os.Getenv("SECRET_KEY")

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	token := make([]byte, length)
	for i := 0; i < length; i++ {
		token[i] = key[r.Intn(len(key))]
	}
	return string(token)
}
func GenerateVerificationCode() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return strconv.Itoa(random.Intn(999999))
}
