package handlers

import (
	"log"
	"project/internal/database"
	"project/internal/logic"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *Manager) LogHandler(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.HTML(200, "login.html", nil)
	case "POST":
		time.Sleep(2 * time.Second) ////// art. delay
		var data database.LoginData

		if err := c.BindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"message": "Неверный формат данных",
			})
			return
		}
		user, boolean, err := m.DB.CheckIfUserExists(data.Email, data.Password)
		if err != nil {
			log.Println(err.Error())
			c.JSON(500, gin.H{
				"message": "Internal server Error",
			})
			return

		}
		if !boolean { // not authorised
			c.JSON(400, gin.H{
				"message": "Пользователя не существует или пароль не верный",
			})
			return
		} else { // authorised
			sessionId, err := logic.GenerateSessionID()
			if err != nil {
				log.Println(err.Error())
				c.JSON(500, gin.H{
					"message": "Internal server error",
				})
				return
			}
			err = m.DB.AuthoriseUserById(user.Id, sessionId)
			if err != nil {
				log.Println(err.Error())
				c.JSON(500, gin.H{
					"message": "Internal server error",
				})
				return
			}
			c.SetCookie("session_id", sessionId, 86400, "/", "", false, false) // name of cookie , sesison id of cookie , max age if cookie , path for cookie , domain  , secure , http only
			c.JSON(200, gin.H{
				"message": "Успешная авторизация",
			})

		}
	}
}
