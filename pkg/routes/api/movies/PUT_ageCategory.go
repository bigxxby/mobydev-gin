package movies

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

func (m *MoviesRoute) PUT_MovieAgeCategory(c *gin.Context) {
	movieId := c.Param("id")
	userRole := c.GetString("role")
	userId := c.GetInt("userId")

	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	valid, movieIdNum := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}
	data := struct {
		AgeCategoryName string `json:"ageCategoryName" binding:"required"`
	}{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	ageCategory, err := m.DB.AgeRepository.GetAgeCategoryByName(data.AgeCategoryName)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "ageCategory not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	err = m.DB.MovieRepository.UpdateMovieAgeCategory(movieIdNum, ageCategory.ID)
	if err != nil {
		log.Println(err.Error())
		c.JSON(404, gin.H{
			"message": "Movie not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Movie ageCategory updated",
	})

}
