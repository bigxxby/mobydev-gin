package routes

import (
	"log"
	"project/internal/utils/mapping"

	"github.com/gin-gonic/gin"
)

func (m *Manager) GET_Movie(c *gin.Context) {
	movies, err := m.DB.GetMovies()
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"movies": mapping.TrimMovies(movies),
	})
}
