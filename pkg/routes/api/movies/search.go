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
	userId := c.GetInt("userId")
	movies, err := m.DB.MovieRepository.SearchMovies(query, userId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"movies": movies})
}

// func (m *MoviesRoute) GET_MostPopular(c *gin.Context) {
// 	movies, err := m.DB.MovieRepository.SearchMovie(query)
// 	if err != nil {
// 		log.Println(err.Error())
// 		c.JSON(500, gin.H{
// 			"message": "Internal server error",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"movies": movies})
// }

// func (m *MoviesRoute) GET_Filters(c *gin.Context) {
// 	// Получение параметров запроса
// 	//filter
// 	genre := c.Query("genre")
// 	category := c.Query("category")
// 	age := c.Query("age")
// 	//order
// 	limit := c.DefaultQuery("limit", "10") // Устанавливаем значение по умолчанию, если параметр отсутствует

// 	newest := c.Query("newest")
// 	oldest := c.Query("oldest")
// 	popular := c.Query("popular")

// }
