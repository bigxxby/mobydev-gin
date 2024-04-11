package routes

import (
	"database/sql"
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
		if err == sql.ErrNoRows {
			log.Println(err.Error())
			c.JSON(404, gin.H{
				"message": "User not found",
			})
			return

		}

		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
	}
	userJson := mapping.TrimUser(*user)
	c.JSON(200, userJson)
}
