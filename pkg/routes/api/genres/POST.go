package genres

import (
	"log"
	"net/http"
	"project/internal/database/genres"

	"github.com/gin-gonic/gin"
)

//	type Category struct {
//		ID          int    `json:"id"`
//		UserID      int    `json:"user_id"` // created by
//		Name        string `json:"category_name"`
//		Description string `json:"description"`
//		Created_at  string `json:"created_at"`
//	}
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

	exists, err := m.DB.GenreRepository.CheckGenreExistsByName(genre.Name)
	if exists {
		c.JSON(409, gin.H{
			"message": "Genre with this name already exists",
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