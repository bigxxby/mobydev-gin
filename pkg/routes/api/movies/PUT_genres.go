package movies

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

func (m *MoviesRoute) PUT_MovieGenres(c *gin.Context) {
	movieId := c.Param("id")
	userRole := c.GetString("role")
	userId := c.GetInt("userId")

	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	valid, movieIdNum := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}
	data := struct {
		Genres []string `json:"genres" binding:"required"`
	}{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	if len(data.Genres) == 0 {
		c.JSON(400, gin.H{
			"message": "At least one genre required",
		})
		return
	}
	genresIdsNew := []int{}
	for _, genre := range data.Genres {
		genreId, _, err := m.DB.GenreRepository.CheckGenreExistsByName(genre)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(404, gin.H{
					"message": "Genre (" + genre + ") does not exists",
				})
				return
			}
			log.Println(err.Error())
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}
		genresIdsNew = append(genresIdsNew, genreId)
	}
	err = m.DB.MovieRepository.AddGenresToMovie(movieIdNum, genresIdsNew)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	genresIdCurrent, err := m.DB.MovieRepository.GetGenresByMovieId(movieIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return

	}
	err = m.DB.MovieRepository.RemoveGenresFromMovie(movieIdNum, genresIdCurrent)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	err = m.DB.MovieRepository.AddGenresToMovie(movieIdNum, genresIdsNew)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Genres Updated",
	})
}
