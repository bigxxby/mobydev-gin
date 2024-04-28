package movies

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *MoviesRoute) GET_Search(c *gin.Context) {
	query := c.Query("query")
	genre := c.Query("genre")
	category := c.Query("category")
	ageCategory := c.Query("ageCategory")
	userId := c.GetInt("userId")
	option := ""
	optionData := ""
	if query != "" {
		movies, err := m.DB.MovieRepository.SearchMovie(userId, query)
		if err != nil {
			log.Println(err.Error())
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"movies": movies})
		return

	} else if genre != "" {
		option = "genre"
		optionData = genre
	} else if category != "" {
		optionData = category
		option = "category"
	} else if ageCategory != "" {
		optionData = ageCategory
		option = "ageCategory"
	}
	movies, err := m.DB.MovieRepository.GetMoviesByFilter(userId, option, optionData)
	if err != nil {
		if err.Error() == "unsupported option" {
			c.JSON(400, gin.H{
				"message": "Bad request",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"movies": movies})
	// if err != nil {
	// 	log.Println(err.Error())
	// 	c.JSON(500, gin.H{
	// 		"message": "Internal server error",
	// 	})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"movies": movies})
	// return
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
