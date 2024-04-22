package auth

import (
	"log"
	"net/http"
	"project/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *AuthRoute) POST_SignIn(c *gin.Context) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	time.Sleep(1 * time.Second) ////// art. delay

	data := struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message": "Неверный формат данных",
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
			"message": "Пользователя не существует или пароль не верный",
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
			"message": "Успешная авторизация",
			"token":   token,
		})

	}
}
func (m *AuthRoute) POST_CheckAuth(c *gin.Context) {
	userRole := c.GetString("role")
	c.JSON(200, gin.H{
		"message": userRole,
	})
}
