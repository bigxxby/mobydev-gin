package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"gopkg.in/mail.v2"
)

func GenerateRandomCode(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)

	var code string
	for i := 0; i < length; i++ {
		code += strconv.Itoa(randGen.Intn(10))
	}
	return code
}
func SendRecoveryCode(email, code string) error {
	from := "test_app@gmail.com"
	password := "your_password"
	to := []string{email}

	// Настройка SMTP клиента
	d := mail.NewDialer("smtp.gmail.com", 587, from, password)

	// Создание письма
	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", "Код восстановления пароля")
	m.SetBody("text/plain", fmt.Sprintf("Ваш код для восстановления пароля: %s", code))

	// Отправка письма
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
