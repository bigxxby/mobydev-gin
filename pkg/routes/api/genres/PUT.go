package genres

import (
	"log"
	"net/http"
	"project/internal/database/genres"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Update a genre
// @Description Updates a genre by its ID if the user has admin privileges
// @Tags genres
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Genre ID"
// @Param genre body genres.GenreShort true "Genre object"
// @Success 200 {object} routes.DefaultMessageResponse "Successful update"
// @Failure 400 {object} routes.DefaultMessageResponse "Bad request"
// @Failure 401 {object} routes.DefaultMessageResponse "Unauthorized"
// @Failure 404 {object} routes.DefaultMessageResponse "Genre Not found"
// @Failure 409 {object} routes.DefaultMessageResponse "Genre with this name already exists"
// @Failure 500 {object} routes.DefaultMessageResponse "Internal server error"
// @Router /api/genres/{id} [put]
func (m *GenreRoute) PUT_Genre(c *gin.Context) {
	genreId := c.Param("id")
	userRole := c.GetString("role")
	userId := c.GetInt("userId")

	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var genre genres.Genre

	valid, genreIdNum := utils.IsValidNum(genreId)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	err := c.BindJSON(&genre)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	exists, _ := m.DB.GenreRepository.CheckGenreExistsById(genreIdNum)
	if !exists {
		c.JSON(404, gin.H{
			"message": "Genre Not found",
		})
		return
	}

	existingGenre, err := m.DB.GenreRepository.GetGenreById(genreIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if existingGenre.Name != genre.Name {
		_, exists, _ = m.DB.GenreRepository.CheckGenreExistsByName(genre.Name)
		if exists {
			c.JSON(409, gin.H{
				"message": "Genre with this name already exists",
			})
			return
		}

	}

	err = m.DB.GenreRepository.UpdateGenre(genreIdNum, genre.Name, genre.Description)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Genre Updated",
	})
}
