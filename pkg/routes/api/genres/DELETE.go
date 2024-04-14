package genres

import (
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

func (m *GenreRoute) DELETE_Genre(c *gin.Context) {
	genreId := c.Param("id")
	userRole := c.GetString("role")
	userId := c.GetInt("userId")

	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	valid, genreIdNum := utils.IsValidNum(genreId)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}

	exists, err := m.DB.GenreRepository.CheckGenreExistsById(genreIdNum)
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

	used, err := m.DB.GenreRepository.CheckGenreIsUsedInMovies(genreIdNum)
	if used {
		log.Println("Can't delete genre when in use of other movies")
		c.JSON(400, gin.H{
			"message": "Cannot delete genre because it is used in movies",
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

	err = m.DB.GenreRepository.DeleteGenreById(genreIdNum)
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
