package users

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils/mapping"

	"github.com/gin-gonic/gin"
)

// GET_Profile retrieves user profile
//	@Summary		Get current users profile
//	@Description	Retrieves the profile of the authenticated user
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{object}	routes.UserProfileResponse		"User profile"
//	@Failure		401	{object}	routes.DefaultMessageResponse	"Unauthorized"
//	@Failure		404	{object}	routes.DefaultMessageResponse	"User not found"
//	@Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
//	@Router			/api/profile [get]
func (m *UsersRoute) GET_Profile(c *gin.Context) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	userId := c.GetInt("userId")
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorised",
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
