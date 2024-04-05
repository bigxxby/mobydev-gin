package routes

import (
	"log"
	"net/http"
	"project/internal/database"
	"project/internal/logic"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *Manager) GET_Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
func (m *Manager) POST_Login(c *gin.Context) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	time.Sleep(1 * time.Second) ////// art. delay
	var data database.LoginData

	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message": "Неверный формат данных",
		})
		return
	}
	user, boolean, err := m.DB.CheckUserExists(data.Email, data.Password)
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
		token, err := logic.CreateJWTToken(user.Id, user.Email, user.Name.String)
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
