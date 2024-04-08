package routes

import (
	"log"
	"net/mail"

	"github.com/gin-gonic/gin"
)

func (m *Manager) GET_Forgot(c *gin.Context) {
	c.HTML(200, "forgot.html", nil)
}
func (m *Manager) POST_Forgot(c *gin.Context) {

	data := struct {
		Email string `json:"email"`
	}{}
	err := c.BindJSON(&data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	_, err = mail.ParseAddress(data.Email)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	// log.Println(data.Email)

	// randomCode := utils.GenerateRandomCode(5)

}
