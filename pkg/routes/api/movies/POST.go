package movies

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/database/movie"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// create movie
func (m *MoviesRoute) POST_Movie(c *gin.Context) {
	userRole := c.GetString("role")
	userId := c.GetInt("userId")

	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	var movie movie.Movie

	err := c.BindJSON(&movie)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	exists, err := m.DB.GenreRepository.CheckGenreExistsById(movie.GenreId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	if !exists {
		c.JSON(400, gin.H{
			"message": "This genre does not exists",
		})
		return
	}
	exists, err = m.DB.AgeRepository.CheckAgeCategoryExistsId(movie.AgeCategoryId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	if !exists {
		c.JSON(400, gin.H{
			"message": "This age category does not exists",
		})
		return
	}
	exists, err = m.DB.CategoriesRepository.CheckCategoryExistsById(movie.CategoryId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return

	}
	if !exists {
		c.JSON(400, gin.H{
			"message": "This category does not exists",
		})
		return
	}

	_, err = m.DB.MovieRepository.CreateMovie(
		userId,
		movie.ImageUrl,
		movie.Name,
		movie.Year,
		movie.CategoryId,
		movie.AgeCategoryId,
		movie.GenreId,
		movie.DurationMinutes,
		movie.Description,
		movie.Keywords,
		movie.Director,
		movie.Producer)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Movie Created",
	})

}

func (m *MoviesRoute) GET_Watch(c *gin.Context) {
	userId := c.GetInt("userId")
	movieId := c.Param("id")
	valid, movieIdNum := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	_, err := m.DB.MovieRepository.GetMovieById(movieIdNum)
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
	count, err := m.DB.MovieRepository.MovieWasWatchedByUser(movieIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": count,
	})

}
