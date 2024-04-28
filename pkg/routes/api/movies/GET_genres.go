package movies

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (m *MoviesRoute) GET_MoviesByGenre(c *gin.Context) {
	genre := c.Query("genre")
	userId := c.GetInt("userId")
	if genre == "" {

		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	movies, err := m.DB.MovieRepository.GetMoviesByGenre(userId, genre)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"movie": movies,
	})
}
