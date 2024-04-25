package movies

import (
	"database/sql"
	"log"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

func (m *MoviesRoute) GET_Movies(c *gin.Context) {
	limit := c.Query("limit")
	if limit == "" {
		movies, err := m.DB.MovieRepository.GetMovies()
		if err != nil {
			log.Println(err.Error())
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}
		c.JSON(200, gin.H{
			"movies": movies,
		})
		return
	}
	valid, num := utils.IsValidNum(limit)
	if !valid {
		log.Println("number not valid")
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	movies, err := m.DB.MovieRepository.GetMoviesLimit(num)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"movies": movies,
	})

}
func (m *MoviesRoute) GET_Movie(c *gin.Context) {
	movieId := c.Param("id")
	valid, num := utils.IsValidNum(movieId)
	if !valid {
		log.Println("num is not valid")
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	movie, err := m.DB.MovieRepository.GetMovieById(num)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Movie not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, movie)
}
