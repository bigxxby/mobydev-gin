package movies

import (
	"log"
	"net/http"
	"project/internal/database/movie"

	"github.com/gin-gonic/gin"
)

// create movie
func (m *MoviesRoute) POST_Movie(c *gin.Context) {
	userRole := c.GetString("role")
	userId := c.GetInt("userId")

	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	var movie movie.Movie

	err := c.BindJSON(&movie)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	_, err = m.DB.MovieRepository.CreateMovie(
		userId,
		movie.ImageUrl,
		movie.Name,
		movie.Year,
		movie.CategoryId,
		movie.AgeCategoryId,
		movie.GenreId,
		movie.DurationMinutes,
		movie.Description,
		movie.Keywords,
		movie.Director,
		movie.Producer)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Movie Created",
	})

}
