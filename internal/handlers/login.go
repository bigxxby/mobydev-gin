package handlers

import "github.com/gin-gonic/gin"

func LogHandler(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
