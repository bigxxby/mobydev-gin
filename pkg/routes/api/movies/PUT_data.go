package movies

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

func (m *MoviesRoute) PUT_MovieData(c *gin.Context) {
	movieId := c.Param("id")
	userRole := c.GetString("role")
	userId := c.GetInt("userId")

	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	valid, mobieIdNum := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}
	data := struct {
		Name            string `json:"name" binding:"required"`
		Description     string `json:"description" binding:"required"`
		Year            int    `json:"year" binding:"required"`
		DurationMinutes int    `json:"durationMinutes" binding:"required"`
		Director        string `json:"director" binding:"required"`
		Producer        string `json:"producer" binding:"required"`
	}{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}
	err = m.DB.MovieRepository.CheckMovieExistsById(mobieIdNum)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Movie not found",
			})
			return
		}
		log.Println(err.Error())
	}
	err = m.DB.MovieRepository.UpdateMovieData(mobieIdNum, data.Name, data.Description, data.Director, data.Producer, data.Year, data.DurationMinutes)
	if err != nil {
		log.Println(err.Error())
		c.JSON(404, gin.H{
			"message": "Movie not found",
		})
		return

	}
	c.JSON(200, gin.H{
		"message": "Movie data updated",
	})

}

// movie := struct {
// 	Name            string   `json:"name" binding:"required"`
// 	Year            int      `json:"year" binding:"required"`
// 	Category        string   `json:"category"`
// 	AgeCategory     string   `json:"ageCategory"`
// 	Genres          []string `json:"genres"` //many to many table
// 	CategoryId      int      `json:"categoryId" binding:"required"`
// 	AgeCategoryId   int      `json:"ageCategoryId" binding:"required"`
// 	DurationMinutes int      `json:"durationMinutes" binding:"required"`
// 	Keywords        string   `json:"keywords" binding:"required"`
// 	Description     string   `json:"description" binding:"required"`
// 	Director        string   `json:"director" binding:"required"`
// 	Producer        string   `json:"producer" binding:"required"`
// }{}
