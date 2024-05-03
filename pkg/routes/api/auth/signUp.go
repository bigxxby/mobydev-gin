package auth

import (
	"log"
	logic "project/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// @Tags			authentication
// @Summary		Sign up
// @Description	Register a new user
// @Accept			json
// @Produce		json
// @Param			body	body		routes.SignUpRequest			true	"Email, password, and confirm password"
// @Success		200		{object}	routes.DefaultMessageResponse	"Successfully registered"
// @Failure		400		{object}	routes.DefaultMessageResponse	"Invalid request format or passwords don't match"
// @Failure		400		{object}	routes.DefaultMessageResponse	"Passwords don't match"
// @Failure		400		{object}	routes.DefaultMessageResponse	"User with specified email already exists"
// @Failure		400		{object}	routes.DefaultMessageResponse	"Invalid email or password format"
// @Failure		500		{object}	routes.DefaultMessageResponse	"Internal server error"
// @Router			/api/signUp [post]
func (m *AuthRoute) POST_SignUp(c *gin.Context) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	time.Sleep(1 * time.Second) //art. delay
	data := struct {
		Email           string `json:"email" binding:"required"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirmPassword" binding:"required"`
	}{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	if data.Password != data.ConfirmPassword {
		c.JSON(400, gin.H{
			"message": "Passwords don't match",
		})
		return
	}

	err := logic.CheckValidForReg(data.Email, data.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	exists, err := m.DB.UserRepository.CreateUser(data.Email, data.Password)
	if exists {
		c.JSON(400, gin.H{
			"message": "User already exists",
		})
		return
	}
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server Error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Registered",
	})

}
