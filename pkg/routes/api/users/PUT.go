package users

import (
	"log"
	"net/http"
	"project/internal/database/user"
	"project/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// PUT_Profile updates current user profile
// @Summary Update current users profile
// @Description Retrieves the profile of the authenticated user
// @Accept json
// @Produce json
// @Param user body routes.UserProfileRequest true "User object to be updated"
// @Security ApiKeyAuth
// @Success 200 {object} routes.DefaultMessageResponse "Profile updated"
// @Failure 400 {object} routes.DefaultMessageResponse "Bad request"
// @Failure 401 {object} routes.DefaultMessageResponse "Unauthorized"
// @Failure 500 {object} routes.DefaultMessageResponse "Internal server error"
// @Router /api/profile [put]
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
	if len(user.Name) > 16 {
		log.Println("Name is too long")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Name is too long",
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
			"message": "Phone Number is not valid",
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

	c.JSON(200, gin.H{
		"message": "Profile updated",
	})
}
