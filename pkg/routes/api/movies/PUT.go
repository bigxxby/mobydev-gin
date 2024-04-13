package movies

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/database/movie"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// update movie

func (m *MoviesRoute) PUT_Movie(c *gin.Context) {
	movieId := c.Param("id")
	userRole := c.GetString("role")
	userId := c.GetInt("userId")

	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	valid, num := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}

	var movie movie.Movie
	err := c.BindJSON(&movie)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	err = m.DB.MovieRepository.CheckMovieExistsById(num)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Movie not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	err = m.DB.MovieRepository.UpdateMovie(
		num, // movie id we want to update
		movie.ImageUrl,
		movie.Name,
		movie.Category,
		movie.MovieType,
		movie.Year,
		movie.AgeCategory,
		movie.DurationMinutes,
		movie.Keywords,
		movie.Description,
		movie.Director,
		movie.Producer,
	)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Movie Updated",
	})

}
