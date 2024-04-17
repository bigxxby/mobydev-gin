package movies

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *MoviesRoute) GET_Search(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	movies, err := m.DB.MovieRepository.SearchMovie(query)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"movies": movies})
}
