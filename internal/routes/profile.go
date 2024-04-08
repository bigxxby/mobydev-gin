package routes

import (
	"log"
	"net/http"
	"project/internal/utils"
	"project/internal/utils/mapping"

	"github.com/gin-gonic/gin"
)

func (m *Manager) GET_Profile(c *gin.Context) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user, err := m.DB.GetUserById(userId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	userJson := mapping.TrimUser(*user)
	c.JSON(200, userJson)
}
