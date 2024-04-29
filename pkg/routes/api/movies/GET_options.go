package movies

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (m *MoviesRoute) GET_EveryMovieByGenre(c *gin.Context) {
	// userId := c.GetInt("userId")
	movies, err := m.DB.MovieRepository.GetEveryMovieByGenre()
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, movies)
}
func (m *MoviesRoute) GET_EveryMovieByCategory(c *gin.Context) {
	movies, err := m.DB.MovieRepository.GetEveryMovieByCategory()
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, movies)
}
func (m *MoviesRoute) GET_EveryMovieByAgeCategory(c *gin.Context) {
	movies, err := m.DB.MovieRepository.GetEveryMovieByAgeCategory()
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, movies)
}
