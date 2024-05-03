package movies

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Tags			movies
//
//	@Summary		Update movie genres
//	@Description	Updates the genres of a movie with the specified ID
//	@Produce		json
//	@Param			id	path	string	true	"Movie ID"
//	@Security		ApiKeyAuth
//	@Param			genres	body		routes.MovieGenresRequest		true	"Genres"
//	@Success		200		{object}	routes.DefaultMessageResponse	"Genres Updated"
//	@Failure		400		{object}	routes.DefaultMessageResponse	"Bad request"
//	@Failure		401		{object}	routes.DefaultMessageResponse	"Unauthorized"
//	@Failure		404		{object}	routes.DefaultMessageResponse	"Genre does not exist"
//	@Failure		500		{object}	routes.DefaultMessageResponse	"Internal server error"
//	@Router			/api/movies/genres/{id} [put]
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
		log.Println(err.Error())
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
	encountered := make(map[string]bool)
	for _, str := range data.Genres {
		if encountered[str] {
			c.JSON(400, gin.H{
				"message": "Genres must be unique",
			})
			return
		}
		encountered[str] = true
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
