package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/database"

	"github.com/gin-gonic/gin"
)

func (m *Manager) Logout(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Method not allowed",
		})
		return
	} else {

		var sessionData database.SessionData
		if err := c.BindJSON(&sessionData); err != nil {
			c.JSON(400, gin.H{
				"message": "Неверный формат данных",
			})
			return
		}
		log.Println(sessionData)

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
}
