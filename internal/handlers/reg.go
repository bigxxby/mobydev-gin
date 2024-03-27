package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RegHandler(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.HTML(200, "reg.html", nil)
	case "POST":
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(500, gin.H{
				"error": "internal server error ",
			})
			return
		}
		log.Println(string(data))
	}
}
