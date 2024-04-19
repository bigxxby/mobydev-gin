package routes

import "github.com/gin-gonic/gin"

func (m *Manager) GET_HTML_Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
func (m *Manager) GET_HTML_Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func (m *Manager) GET_HTML_Reg(c *gin.Context) {

	c.HTML(200, "reg.html", nil)
}
func (m *Manager) GET_Restore(c *gin.Context) {
	c.HTML(200, "restore.html", nil)
}
func (m *Manager) GET_HTML_Movie(c *gin.Context) {
	c.HTML(200, "movie_create.html", nil)
}
