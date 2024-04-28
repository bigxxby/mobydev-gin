package movies

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// PUT_MovieData updates the data of a movie
// @Summary Update movie data
// @Description Updates the data of a movie with the specified ID
// @Produce json
// @Param id path string true "Movie ID"
// @Security ApiKeyAuth
// @Param movieData body routes.MovieDataRequest true "Movie Data"
// @Success 200 {object} routes.DefaultMessageResponse "Movie data updated"
// @Failure 400 {object} routes.DefaultMessageResponse "Bad request"
// @Failure 401 {object} routes.DefaultMessageResponse "Unauthorized"
// @Failure 404 {object} routes.DefaultMessageResponse "Movie not found"
// @Failure 500 {object} routes.DefaultMessageResponse "Internal server error"
// @Router /api/movies/data/{id} [put]
func (m *MoviesRoute) PUT_MovieData(c *gin.Context) {
	movieId := c.Param("id")
	userRole := c.GetString("role")
	userId := c.GetInt("userId")

	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	valid, mobieIdNum := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	data := struct {
		Name            string `json:"name" binding:"required"`
		Description     string `json:"description" binding:"required"`
		Year            int    `json:"year" binding:"required"`
		DurationMinutes int    `json:"durationMinutes" binding:"required"`
		Director        string `json:"director" binding:"required"`
		Producer        string `json:"producer" binding:"required"`
		Keywords        string `json:"keywords" binding:"required"`
	}{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	err = m.DB.MovieRepository.CheckMovieExistsById(mobieIdNum)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Movie not found",
			})
			return
		}
		log.Println(err.Error())
	}
	err = m.DB.MovieRepository.UpdateMovieData(mobieIdNum, data.Name, data.Description, data.Director, data.Producer, data.Year, data.DurationMinutes, data.Keywords)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return

	}
	c.JSON(200, gin.H{
		"message": "Movie data updated",
	})

}
