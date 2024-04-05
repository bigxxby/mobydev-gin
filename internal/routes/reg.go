package routes

import (
	"log"
	"project/internal/database"
	logic "project/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *Manager) GET_Reg(c *gin.Context) {

	c.HTML(200, "reg.html", nil)
}
func (m *Manager) POST_Reg(c *gin.Context) {

	time.Sleep(1 * time.Second) //art. delay
	var data database.RegisterData

	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message": "Неверный формат данных",
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
	exists, err := m.DB.CreateUser(data.Email, data.Password)
	if exists {
		c.JSON(400, gin.H{
			"message": "Пользователь уже существует",
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
		"message": "Успешно зарегистрирован",
	})

}
