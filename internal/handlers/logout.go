package handlers

import (
	"database/sql"
	"log"
	"project/internal/database"

	"github.com/gin-gonic/gin"
)

// /api/logout
func (m *Manager) POST_Logout(c *gin.Context) {
	var sessionData database.SessionData
	if err := c.BindJSON(&sessionData); err != nil {
		c.JSON(400, gin.H{
			"message": "Неверный формат данных",
		})
		return
	}

	if sessionData.SessionID == "" {
		c.JSON(400, gin.H{
			"message": "SessionId cant be empty string",
		})
		return
	}

	err := m.DB.LogoutUser(sessionData.SessionID)
	if err == sql.ErrNoRows {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "No user with this sesssionId is found",
		})
		return
	}
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server Error",
		})
		log.Println(err.Error())
		return
	}
	c.SetCookie("session_id", "", 1, "/", "", false, false) // name of cookie , sesison id of cookie , max age if cookie , path for cookie , domain  , secure , http only

	c.JSON(200, nil)
}
