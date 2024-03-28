package handlers

import (
	"project/internal/database"
	"project/internal/logic"

	"github.com/gin-gonic/gin"
)

func (m *Manager) RegHandler(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.HTML(200, "reg.html", nil)
	case "POST":
		var data database.RegisterData

		if err := c.BindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid JSON payload",
			})
			return
		}

		err := logic.CheckValidForReg(data.Email, data.Password, data.ConfirmPassword)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		err = m.DB.AddUser(data.Email, data.Password)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal server Error",
			})
			return
		}
		// После успешного добавления пользователя, вы можете перенаправить пользователя на другую страницу или отправить ответ об успешной регистрации
		c.JSON(200, gin.H{
			"message": "Registration successful",
		})
	}
}
