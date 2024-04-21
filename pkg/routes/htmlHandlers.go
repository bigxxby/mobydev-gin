package routes

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

func (m *Manager) GET_HTML_Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
func (m *Manager) GET_HTML_Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func (m *Manager) GET_HTML_Reg(c *gin.Context) {

	c.HTML(200, "reg.html", nil)
}
func (m *Manager) GET_HTML_SendRestoreCode(c *gin.Context) {
	c.HTML(200, "sendCode.html", nil)
}
func (m *Manager) GET_HTML_Movie(c *gin.Context) {
	c.HTML(200, "movie_create.html", nil)
}

// gives permission to change password HTML
func (m *Manager) GET_ChangePassword(c *gin.Context) {
	token := c.Query("token")
	email := c.Query("email")
	if token == "" || email == "" {
		log.Println("token or email is null")
		c.JSON(404, gin.H{
			"message": "Not found",
		})
		return
	}
	err := m.DB.UserRepository.VerifyToken(token)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("token or email is null")
			c.JSON(404, gin.H{
				"message": "Not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return

	}
	data := struct {
		Token string
		Email string
	}{
		Token: token,
		Email: email,
	}
	c.HTML(200, "changePass.html", data)
}
