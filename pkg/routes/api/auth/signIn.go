package auth

import (
	"log"
	"net/http"
	"project/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// @Tags authentication
// @Summary		Sign in
//
// @Description	Sign in with email and password to obtain JWT token. Creditnails for sign in "email": "big@example.com",  "password": "Aa12345678#"
//
// @Accept			json
// @Produce		json
// @Param			body	body		routes.SignInRequest			true	"Email and password"
// @Success		200		{object}	routes.SignInResponse			"Successfully signed in"
// @Failure		400		{object}	routes.DefaultMessageResponse	"Invalid data format"
// @Failure		401		{object}	routes.DefaultMessageResponse	"User does not exist or incorrect password"
// @Failure		500		{object}	routes.DefaultMessageResponse	"Internal server error"
// @Router			/api/signIn [post]
func (m *AuthRoute) POST_SignIn(c *gin.Context) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	time.Sleep(1 * time.Second) ////// art. delay
	data := struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	if err := c.BindJSON(&data); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	err := utils.CheckValidForReg(data.Email, data.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	user, boolean, err := m.DB.UserRepository.CheckUserСredentials(data.Email, data.Password)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server Error",
		})
		return
	}
	if !boolean { // not authorised
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found or password incorrect",
		})
		return
	} else { // authorised
		token, err := utils.CreateJWTToken(user.Id, user.Email, user.Name.String, user.Role)
		if err != nil {
			log.Println(err.Error())
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Authorized",
			"token":   token,
		})

	}
}

// @Tags authentication
// @Summary		Check authentication
// @Description	Check user authentication and return user role
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Success		200	{object}	routes.DefaultMessageResponse	"Successfully checked authentication"
// @Failure		401	{object}	routes.DefaultMessageResponse	"Unauthorized: User not authenticated"
// @Router			/api/check-auth [post]
func (m *AuthRoute) POST_CheckAuth(c *gin.Context) {
	userRole := c.GetString("role")
	if userRole == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorised",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": userRole,
	})
}
