package users

import (
	"log"
	"net/http"
	"project/internal/database/user"
	"project/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *UsersRoute) PUT_Profile(c *gin.Context) {
	userId := c.GetInt("userId")
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorised",
		})
		return
	}

	var user user.UserJson
	err := c.BindJSON(&user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	if len(user.Name) < 16 {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	var date time.Time
	if user.DateOfBirth != "" {
		date, err = time.Parse("2006-01-02", user.DateOfBirth)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
			})
			return
		}
	}
	if (user.Phone != "") && !(utils.IsValidPhoneNumber(user.Phone)) {
		log.Println("Phone Number is not valid")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	err = m.DB.UserRepository.UpdateProfile(userId, user.Name, user.Phone, &date)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Profile updated",
	})
}
