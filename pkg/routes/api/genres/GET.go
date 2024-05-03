package genres

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Get all genres
// @Description Retrieves all genres
// @Tags genres
// @Produce json
// @Success 200 {object} []genres.Genre "Successful response"
// @Failure 404 {object} routes.DefaultMessageResponse "No genres added"
// @Failure 500 {object} routes.DefaultMessageResponse "Internal server error"
// @Router /api/genres [get]
func (m *GenreRoute) GET_Genres(c *gin.Context) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	genres, err := m.DB.GenreRepository.GetAllGenres()
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "No genres added",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, genres)
}

// @Summary Get a genre by ID
// @Description Retrieves a genre by its ID
// @Tags genres
// @Produce json
// @Param id path int true "Genre ID"
// @Success 200 {object} genres.Genre "Successful response"
// @Failure 400 {object} routes.DefaultMessageResponse "Bad request"
// @Failure 404 {object} routes.DefaultMessageResponse "Genre not found"
// @Failure 500 {object} routes.DefaultMessageResponse "Internal server error"
// @Router /api/genres/{id} [get]
func (m *GenreRoute) GET_Genre(c *gin.Context) {
	genresId := c.Param("id")
	valid, genresIdNums := utils.IsValidNum(genresId)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}
	genre, err := m.DB.GenreRepository.GetGenreById(genresIdNums)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Genre not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, genre)
}
