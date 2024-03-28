package handlers

import "github.com/gin-gonic/gin"

func (m *Manager) LogHandler(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
