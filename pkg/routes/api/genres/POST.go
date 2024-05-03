package genres

import (
	"log"
	"net/http"
	"project/internal/database/genres"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new genre
// @Description Creates a new genre if the user has admin privileges
// @Tags genres
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param genre body genres.GenreShort true "Genre object"
// @Success 200 {object} routes.DefaultMessageResponse "Successful creation"
// @Failure 400 {object} routes.DefaultMessageResponse "Bad request"
// @Failure 401 {object} routes.DefaultMessageResponse "Unauthorized"
// @Failure 409 {object} routes.DefaultMessageResponse "Genre with this name already exists"
// @Failure 500 {object} routes.DefaultMessageResponse "Internal server error"
// @Router /api/genres [post]
func (m *GenreRoute) POST_Genre(c *gin.Context) {
	userRole := c.GetString("role")
	userId := c.GetInt("userId")

	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	var genre genres.Genre

	err := c.BindJSON(&genre)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	_, exists, _ := m.DB.GenreRepository.CheckGenreExistsByName(genre.Name)
	if exists {
		c.JSON(409, gin.H{
			"message": "Genre with this name already exists",
		})
		return
	}

	_, err = m.DB.GenreRepository.CreateGenre(userId, genre.Name, genre.Description)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Genre Created",
	})

}
