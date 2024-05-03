package genres

import (
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Delete a genre
// @Description Deletes a genre by its ID if the user has admin privileges
// @Tags genres
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Genre ID"
// @Success 200 {object} routes.DefaultMessageResponse "Successful deletion"
// @Failure 400 {object} routes.DefaultMessageResponse "Bad request"
// @Failure 401 {object} routes.DefaultMessageResponse "Unauthorized"
// @Failure 404 {object} routes.DefaultMessageResponse "Genre not found"
// @Failure 409 {object} routes.DefaultMessageResponse "Genre is used in other movies"
// @Failure 500 {object} routes.DefaultMessageResponse "Internal server error"
// @Router /api/genres/{id} [delete]
func (m *GenreRoute) DELETE_Genre(c *gin.Context) {
	userRole := c.GetString("role")
	userId := c.GetInt("userId")
	genreId := c.Param("id")
	valid, genresIdNum := utils.IsValidNum(genreId)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}
	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	exists, err := m.DB.GenreRepository.CheckGenreExistsById(genresIdNum)
	if !exists {
		c.JSON(404, gin.H{
			"message": "Genre not found",
		})
		return
	}
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	used, err := m.DB.GenreRepository.CheckGenreIsUsedInMovies(genresIdNum)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	if used {
		c.JSON(409, gin.H{
			"message": "Genre is used in other movies",
		})
		return
	}

	err = m.DB.GenreRepository.DeleteGenreById(genresIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Genre Deleted",
	})

}
