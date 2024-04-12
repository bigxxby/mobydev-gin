package routes

import (
	"database/sql"
	"log"
	"project/internal/utils/mapping"

	"github.com/gin-gonic/gin"
)

func (m *Manager) GET_Profile(c *gin.Context) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	userId := c.GetInt("userId")
	if userId == 0 {
		log.Println("middleware error")
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	user, err := m.DB.UserRepository.GetUserById(userId)
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
