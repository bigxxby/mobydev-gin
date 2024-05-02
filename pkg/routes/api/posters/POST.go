package posters

import (
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// adds posters to he movie by array of 5 (CAN ACCEPT only 5 ulrs)
// @Tags Posters
func (m *PosterRoute) POST_PostersOfMoive(c *gin.Context) {
	movieId := c.Param("id")
	userId := c.GetInt("userId")
	userRole := c.GetString("role")
	if userId == 0 || userRole != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"messsage": "Unauthorized",
		})
		return
	}
	valid, movieIdNum := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
		})
		return
	}
	data := struct {
		Posters [5]string `json:"posters" binding:"required"`
	}{}
	err := c.BindJSON(&data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
		})
		return
	}

	err = m.DB.MovieRepository.CheckMovieExistsById(movieIdNum)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Movie not found",
		})
		return
	}
	has, err := m.DB.PosterRepo.CheckIfMovieHaveNoPosters(movieIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if has {
		err = m.DB.PosterRepo.DeletePostersOfMovie(movieIdNum)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}
	}
	err = m.DB.PosterRepo.AddPostersMovieById(movieIdNum, data.Posters)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Posters Added",
	})
}
