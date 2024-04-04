package handlers

import (
	"log"
	"net/http"
	"project/internal/database"

	"github.com/gin-gonic/gin"
)

func (m *Manager) GET_Profile(c *gin.Context) {
	sessionId := c.Param("sessionId")
	if sessionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Session not found",
		})
		return
	}
	user, err := m.DB.FindUserBySessionId(sessionId)
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
