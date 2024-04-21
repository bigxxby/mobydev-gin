package auth

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	m2 "net/mail"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

func (m *AuthRoute) POST_VerifyCode(c *gin.Context) {

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

	token := utils.GenerateResetToken(20)

	err = m.DB.UserRepository.SaveResetToken(token, data.Email)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	err = utils.SendTokenEmail(data.Email, token)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Failed to send email",
		})
		return
	}
	redirectStr := fmt.Sprintf("http://localhost:8080/change-password?token=%s&email=%s", token, data.Email)
	c.JSON(http.StatusPermanentRedirect, gin.H{
		"message":  "Code is verified ",
		"redirect": redirectStr,
	})

}
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
	err = utils.CheckValidForReg(data.Email, data.NewPassword, "user")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Password not valid",
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
			"message": "Internal server error",
		})
		return
	}
	err = m.DB.UserRepository.UpdatePassword(data.Email, hash)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	err = m.DB.UserRepository.DeleteResetData(data.Email)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Password updated successfully",
	})
}
func (m *AuthRoute) POST_SendCode(c *gin.Context) {
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
	code := utils.GenerateVerificationCode()
	err = m.DB.UserRepository.AddVerificationCode(data.Email, code)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	err = utils.SendCodeEmail(data.Email, code)
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
