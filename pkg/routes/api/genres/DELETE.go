package genres

import (
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

//	@Tags	Category
//
//	type Category struct {
//		ID          int    `json:"id"`
//		UserID      int    `json:"user_id"` // created by
//		Name        string `json:"category_name"`
//		Description string `json:"description"`
//		Created_at  string `json:"created_at"`
//	}
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
