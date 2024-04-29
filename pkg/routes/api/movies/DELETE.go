package movies

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// DELETE_Movie deletes a movie
//	@Summary		Delete a movie
//	@Description	Deletes a movie with the specified ID
//	@Produce		json
//	@Param			id	path	string	true	"Movie ID"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	routes.DefaultMessageResponse	"Movie Deleted"
//	@Failure		400	{object}	routes.DefaultMessageResponse	"Bad request"
//	@Failure		401	{object}	routes.DefaultMessageResponse	"Unauthorized"
//	@Failure		404	{object}	routes.DefaultMessageResponse	"Movie not found"
//	@Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
//	@Router			/api/movies/{id} [delete]
func (m *MoviesRoute) DELETE_Movie(c *gin.Context) {
	movieId := c.Param("id")
	userRole := c.GetString("role")
	userId := c.GetInt("userId")

	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	valid, num := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}

	err := m.DB.MovieRepository.CheckMovieExistsById(num)
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

	err = m.DB.MovieRepository.DeleteMovie(movieId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Movie Deleted",
	})

}
