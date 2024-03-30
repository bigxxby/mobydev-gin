package handlers

import (
	"log"
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
	user, err := m.DB.FindUserBySessionId(sessionData.SessionID)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server Error",
		})
		log.Println(err.Error())
		return
	}
	userJson := sendBackUser(*user)
	c.JSON(200, userJson)
}

// query := "SELECT id , email , name , phone,date_of_birth FROM users WHERE session_id = $1"
func sendBackUser(user database.User) gin.H {
	userJson := gin.H{}

	userJson["id"] = user.Id
	userJson["email"] = user.Email
	if user.Name.Valid {
		userJson["name"] = user.Name.String
	}
	if user.Phone.Valid {
		userJson["phone"] = user.Phone.String
	}
	if user.DateOfBirth.Valid {
		userJson["dot"] = user.DateOfBirth.Time.String()
	}

	return userJson
}
