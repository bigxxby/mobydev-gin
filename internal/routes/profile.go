package routes

import (
	"log"
	"net/http"
	"project/internal/database"
	"project/internal/utils"

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
	}
	userJson := SendBackUserTrim(*user)
	c.JSON(200, userJson)
}

// query := "SELECT id , email , name , phone,date_of_birth FROM users WHERE session_id = $1"
func SendBackUserTrim(user database.User) gin.H {
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
