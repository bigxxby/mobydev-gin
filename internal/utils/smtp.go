package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendCodeEmail(email, code string) error {
	 from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	to := []string{email}
	subject := "Код подтверждения"
	body := fmt.Sprintf("Ваш код подтверждения: %s", code)

	smtpServer := "smtp.gmail.com"
	smtpPort :=   "587"

	message :=   []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", email, subject, body))

	auth := smtp.PlainAuth("", from, password, smtpServer)
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}
	return nil
}
func SendTokenEmail(email, token string) error {
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	to := []string{email}
	subject := "Временная ссылка для смены пароля"
	body := fmt.Sprintf("Ваша ссылка для смены пароля: http://localhost:8080/restore?token=%s&email=%s", token, email)

	smtpServer := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", email, subject, body))

	auth := smtp.PlainAuth("", from, password, smtpServer)
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}

	return nil
}
