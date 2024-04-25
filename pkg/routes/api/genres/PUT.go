package genres

import (
	"log"
	"net/http"
	"project/internal/database/genres"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// type Genre struct {
// 	ID          int    `json:"id"`
// 	UserID      int    `json:"user_id"` // created by
// 	Name        string `json:"genre_name"`
// 	Description string `json:"description"`
// 	Created_at  string `json:"created_at"`
// }

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

	exists, err := m.DB.GenreRepository.CheckGenreExistsById(genreIdNum)
	if !exists {
		c.JSON(404, gin.H{
			"message": "Genre Not found",
		})
		return
	}
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Internal server error",
		})
		return
	}

	existingGenre, err := m.DB.GenreRepository.GetGenreById(genreIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if existingGenre.Name != genre.Name {
		_, exists, err = m.DB.GenreRepository.CheckGenreExistsByName(genre.Name)
		if exists {
			c.JSON(409, gin.H{
				"message": "Genre with this name already exists",
			})
			return
		}
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Internal server error",
			})
			return
		}
	}

	err = m.DB.GenreRepository.UpdateGenre(genreIdNum, genre.Name, genre.Description)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Genre Updated",
	})
}
