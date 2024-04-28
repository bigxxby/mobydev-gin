package movies

import (
	"log"

	"github.com/gin-gonic/gin"
)

// GET_MoviesByGenre retrieves movies by genre for a user
// @Summary Get movies by genre
// @Description Retrieves movies by genre for a user
// @Produce json
// @Param genre query string true "Genre"
// @Security ApiKeyAuth
// @Success 200 {object} routes.DefaultMessageResponse "OK"
// @Failure 400 {object} routes.DefaultMessageResponse "Bad request"
// @Failure 500 {object} routes.DefaultMessageResponse "Internal server error"
// @Router /api/movies/genre [get]
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
