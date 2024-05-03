package movies

import (
	"database/sql"
	"log"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Tags			movies
// @Summary		Get all movies
// @Description	Retrieves movies for a user
// @Produce		json
// @Param			limit	query		int							false	"Limit"
// @Success		200		{object}	routes.ManyMoviesResponse	"OK"
// @Failure		400		{object}	routes.DefaultMessageResponse"Invalid limit parameter"
// @Failure		500		{object}	routes.DefaultMessageResponse	"Internal server error"
// @Router			/api/movies [get]
func (m *MoviesRoute) GET_Movies(c *gin.Context) {
	userId := c.GetInt("userId")
	limit := c.Query("limit")
	if limit == "" {
		movies, err := m.DB.MovieRepository.GetMovies(userId)
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
	valid, num := utils.IsValidNum(limit)
	if !valid {
		log.Println("number not valid")
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	movies, err := m.DB.MovieRepository.GetMoviesLimit(num, userId)
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

// @Tags			movies
// @Summary		Get movie details
// @Description	Retrieves details of a specific movie including its seasons and similar movies
// @Produce		json
// @Param			id	path		string							true	"Movie ID"
// @Success		200	{object}	routes.MovieResponse			"OK"
// @Failure		400	{object}	routes.DefaultMessageResponse	"Invalid movie ID"
// @Failure		404	{object}	routes.DefaultMessageResponse	"Movie not found"
// @Failure		500	{object}	routes.DefaultMessageResponse"Internal server error"
// @Router			/api/movies/{id} [get]
func (m *MoviesRoute) GET_Movie(c *gin.Context) {
	movieId := c.Param("id")
	userId := c.GetInt("userId")
	valid, movieIdNum := utils.IsValidNum(movieId)
	if !valid {
		log.Println("num is not valid")
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	movie, err := m.DB.MovieRepository.GetMovieById(userId, movieIdNum)
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
	seasons, err := m.DB.SeasonRepository.GetAllSeasonsOfMovieId(movie.Id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	//right not just give 5 random movies
	similar, err := m.DB.MovieRepository.GetSimularMoviesLimit5(movie.Id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"movie":   movie,
		"seasons": seasons,
		"simular": similar,
	})
}
