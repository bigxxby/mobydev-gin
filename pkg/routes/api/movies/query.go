package movies

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary		Search movies
//	@Description	Search movies based on a query string, (searches by name and keywords)
//	@Tags			query
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			query	query		string							true	"Search query"
//	@Success		200		{object}	routes.ManyMoviesResponse		"Successful response"
//	@Failure		500		{object}	routes.DefaultMessageResponse	"Internal server error"
//	@Router			/api/movies/search [get]
func (m *MoviesRoute) GET_Search(c *gin.Context) {
	userId := c.GetInt("userId")
	query := c.Query("query")
	movies, err := m.DB.MovieRepository.SearchMovie(userId, query)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"movies": movies})
}

//	@Summary		Get every movie by genre
//	@Description	Retrieves every movie categorized by genre
//	@Tags			query
//	@Produce		json
//	@Success		200	{object}	map[string][]movie.MovieShort	"Successful response"
//	@Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
//	@Router			/api/movies/by-genre [get]
func (m *MoviesRoute) GET_EveryMovieByGenre(c *gin.Context) {
	// userId := c.GetInt("userId")
	movies, err := m.DB.MovieRepository.GetEveryMovieByGenre()
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, movies)
}

//	@Summary		Get every movie by category
//	@Description	Retrieves every movie categorized by category
//	@Tags			query
//	@Produce		json
//	@Success		200	{object}	map[string][]movie.MovieShort	"Successful response"
//	@Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
//	@Router			/api/movies/by-category [get]
func (m *MoviesRoute) GET_EveryMovieByCategory(c *gin.Context) {
	movies, err := m.DB.MovieRepository.GetEveryMovieByCategory()
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, movies)
}

//	@Summary		Get every movie by age category
//	@Description	Retrieves every movie categorized by age category
//	@Tags			query
//	@Produce		json
//	@Success		200	{object}	map[string][]movie.MovieShort	"Successful response"
//	@Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
//	@Router			/api/movies/by-age-category [get]
func (m *MoviesRoute) GET_EveryMovieByAgeCategory(c *gin.Context) {
	movies, err := m.DB.MovieRepository.GetEveryMovieByAgeCategory()
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, movies)
}
