package handlers

import (
	"log"
	"project/internal/database"
	"project/internal/logic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (m *Manager) GET_Reg(c *gin.Context) {

	c.HTML(200, "reg.html", nil)
}
func (m *Manager) POST_Reg(c *gin.Context) {

	time.Sleep(2 * time.Second) //art. delay
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
	err = m.DB.AddUser(data.Email, data.Password)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				c.JSON(400, gin.H{
					"message": "Пользователь уже существует",
				})
				return
			}
		}
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
