package auth

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	m2 "net/mail"
	"net/smtp"
	"os"
	"project/internal/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *AuthRoute) POST_Restore(c *gin.Context) {

	data := struct {
		Email string `json:"email"`
	}{}

	err := c.BindJSON(&data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	_, err = m2.ParseAddress(data.Email) //is Really email
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	err = m.DB.UserRepository.CheckUserExistsByEmail(data.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "User not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	code := generateVerificationCode()
	err = m.DB.UserRepository.AddVerificationCode(data.Email, code)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	err = sendCodeEmail(data.Email, code)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	log.Println(code)
	c.JSON(200, gin.H{
		"message": "Code sent to the email (check spam too)",
	})

}
func (m *AuthRoute) POST_Verify(c *gin.Context) {

	data := struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}{}

	err := c.BindJSON(&data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	_, err = m2.ParseAddress(data.Email) //is Really email
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	err = m.DB.UserRepository.ValidateCode(data.Email, data.Code)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(400, gin.H{
				"message": "Code is not valid",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	token := generateResetToken(10)

	err = m.DB.UserRepository.SaveResetToken(data.Email, token)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	err = sendTokenEmail(data.Email, token)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Failed to send email",
		})
		return
	}
	redirectStr := fmt.Sprintf("http://localhost:8080/restore?token=%s&email=%s", token, data.Email)
	c.JSON(http.StatusPermanentRedirect, gin.H{
		"message":  "Code is verified ",
		"redirect": redirectStr,
	})

}

func generateResetToken(length int) string {
	key := os.Getenv("SECRET_KEY")

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	token := make([]byte, length)
	for i := 0; i < length; i++ {
		token[i] = key[r.Intn(len(key))]
	}
	return string(token)
}

func (m *AuthRoute) GET_ResetPassword(c *gin.Context) {
	token := c.Query("token")
	email := c.Query("email")
	if token == "" || email == "" {
		log.Println("token or email is null")
		c.JSON(404, gin.H{
			"message": "Not found",
		})
		return
	}
	err := m.DB.UserRepository.VerifyToken(token)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("token or email is null")
			c.JSON(404, gin.H{
				"message": "Not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return

	}
	data := struct {
		Token string
		Email string
	}{
		Token: token,
		Email: email,
	}
	c.HTML(200, "changePass.html", data)

}

// api/reset-password
func (m *AuthRoute) POST_ResetPassword(c *gin.Context) {
	data := struct {
		Email       string `json:"email"`
		ResetToken  string `json:"token"`
		NewPassword string `json:"password"`
	}{}

	err := c.BindJSON(&data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	err = m.DB.UserRepository.VerifyToken(data.ResetToken)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	hash, err := utils.HashPassword(data.NewPassword)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Failed to update password",
		})
		return
	}
	err = m.DB.UserRepository.UpdatePassword(data.Email, hash)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Failed to update password",
		})
		return
	}

	// err = m.DB.UserRepository.DeleteResetToken(data.Email, data.ResetToken)
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	c.JSON(200, gin.H{
		"message": "Password updated successfully",
	})
}

func generateVerificationCode() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return strconv.Itoa(random.Intn(999999))
}

func sendCodeEmail(email, code string) error {
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	to := []string{email}
	subject := "Код подтверждения"
	body := fmt.Sprintf("Ваш код подтверждения: %s", code)

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
func sendTokenEmail(email, token string) error {
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	to := []string{email}
	subject := "Временная ссылка для смены пароля"
	body := fmt.Sprintf("Ваша ссылка для смены пароля: http://localhost:8080/restore?token=%s&email=%s", token, email)

	smtpServer := "smtp.gmail.com"
	smtpPort := "645"

	message := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", email, subject, body))

	auth := smtp.PlainAuth("", from, password, smtpServer)
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}

	return nil
}
