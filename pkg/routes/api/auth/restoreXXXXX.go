package auth

import (
	"log"
	"net/http"
	m2 "net/mail"

	"github.com/gin-gonic/gin"
)

func POST_Restore(c *gin.Context) {

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
	_, err = m2.ParseAddress(data.Email)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

}
