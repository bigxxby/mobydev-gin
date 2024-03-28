package handlers

import (
	"github.com/gin-gonic/gin"
)

func (m *Manager) Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
