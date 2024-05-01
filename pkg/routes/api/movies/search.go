package movies

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

// func (m *MoviesRoute) GET_Filter(c *gin.Context) {
// 	userId := c.GetInt("userId")
// 	genre := c.Query("genre")
// 	category := c.Query("category")
// 	ageCategory := c.Query("ageCategory")

// 	optionData := ""
// 	option := ""
// 	if genre != "" {
// 		optionData = genre
// 		option = "genre"
// 	} else if category != "" {
// 		optionData = category
// 		option = "category"
// 	} else if ageCategory != "" {
// 		optionData = ageCategory
// 		option = "ageCategory"
// 	} else {
// 		c.JSON(400, gin.H{
// 			"message": "Bad request",
// 		})
// 		return
// 	}

// 	movies, err := m.DB.MovieRepository.GetMoviesByFilter(userId, option, optionData)
// 	if err != nil {
// 		log.Println(err.Error())
// 		c.JSON(500, gin.H{
// 			"message": "Internal server error",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"movies": movies})
// }
