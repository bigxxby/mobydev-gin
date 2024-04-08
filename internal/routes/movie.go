package routes

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (m *Manager) GET_Movie(c *gin.Context) {
	limit := c.Query("limit")
	num := 0
	if limit != "" {
		var err error
		num, err = strconv.Atoi(limit)
		if err != nil {
			log.Println(err.Error())
			c.JSON(400, gin.H{
				"message": "Bad request",
			})
			return
		} else if num < 0 {
			log.Println("Limit cant be < 0")
			c.JSON(400, gin.H{
				"message": "Bad request",
			})
			return
		}

	}
	movies, err := m.DB.GetMovies(num)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		// "movies": mapping.TrimMovies(movies),
		"movies": movies,
	})
}
