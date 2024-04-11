package routes

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/database"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

func (m *Manager) GET_HTML_Movie(c *gin.Context) {
	c.HTML(200, "movie_create.html", nil)
}

// get movies
func (m *Manager) GET_Movies(c *gin.Context) {
	limit := c.Query("limit")
	if limit != "" {
		valid, num := utils.IsValidNum(limit)
		if !valid {
			log.Println("number not valid")
			c.JSON(400, gin.H{
				"message": "Bad request",
			})
			return
		}
		movies, err := m.DB.GetMovies(num)
		if err != nil {
			log.Println(err.Error())
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}
		c.JSON(200, gin.H{
			"movies": movies,
		})
		return
	}

	movies, err := m.DB.GetMovies(0)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"movies": movies,
	})
}

// get movie by id
func (m *Manager) GET_Movie(c *gin.Context) {
	movieId := c.Param("id")
	valid, num := utils.IsValidNum(movieId)
	if !valid {
		log.Println("num is not valid")
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	movie, err := m.DB.GetMovieById(num)
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
	c.JSON(200, movie)
}

// create movie
func (m *Manager) POST_Movie(c *gin.Context) {
	token := c.GetHeader("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user, err := m.DB.GetUserById(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "User not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return

	}

	if user.Role != "admin" {
		log.Println("this user is not admin")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	var movie database.Movie

	err = c.BindJSON(&movie)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	_, err = m.DB.CreateMovie(
		userId,
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
		movie.Producer)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Movie Created",
	})

}

// delete movie
func (m *Manager) DELETE_Movie(c *gin.Context) {
	token := c.GetHeader("Authorization")
	movieId := c.Param("id")

	valid, num := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user, err := m.DB.GetUserById(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "User not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return

	}
	if user.Role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	exists, err := m.DB.CheckMovieExistsById(num)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	if !exists {
		log.Println("movie not found")
		c.JSON(404, gin.H{
			"message": "Movie not found",
		})
		return
	}
	err = m.DB.DeleteMovie(movieId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Movie Deleted",
	})

}

// update movie
func (m *Manager) PUT_Movie(c *gin.Context) {
	token := c.GetHeader("Authorization")
	movieId := c.Param("id")

	valid, num := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	user, err := m.DB.GetUserById(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "User not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return

	}
	if user.Role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	var movie database.Movie
	err = c.BindJSON(&movie)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	exists, err := m.DB.CheckMovieExistsById(num)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	if !exists {
		log.Println("movie not found")
		c.JSON(404, gin.H{
			"message": "Movie not found",
		})
		return
	}

	err = m.DB.UpdateMovie(
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
