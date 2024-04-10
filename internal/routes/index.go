package routes

import (
	"github.com/gin-gonic/gin"
)

func (m *Manager) GET_HTML_Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
