package handlers

import (
	"net/http"
	"project/internal/database"

	"github.com/gin-gonic/gin"
)

func (m *Manager) GetProfile(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Method not allowed",
		})
		return
	}
	var sessionData database.SessionData

	if err := c.BindJSON(&sessionData); err != nil {
		c.JSON(400, gin.H{
			"error": "Невозможно привязать JSON",
		})
		return
	}

	// user, err := m.DB.FindUserBySessionId(sessionData.SessionID)
	// if err != nil {
	// c.JSON(500, gin.H{
	// "message": "Internal server Error",
	// })
	// return
	// }
	// log.Println(user)
	c.JSON(200, gin.H{
		"message": "Данные успешно получены",
	})
}
