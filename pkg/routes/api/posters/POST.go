package posters

import (
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

func (m *PosterRoute) POST_PostersOfMoive(c *gin.Context) {
	movieId := c.Param("id")
	valid, movieIdNum := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
		})
		return
	}
	data := struct {
		Id           int    `json:"id"`
		MovieId      int    `json:"movieId"`
		MainPoster   string `json:"mainPoster" binding:"required"`
		SecondPoster string `json:"secondPoster"`
		ThirdPoster  string `json:"thirdPoster"`
		FourthPoster string `json:"fourthPoster"`
		FifthPoster  string `json:"fifthPoster"`
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
		c.JSON(http.StatusConflict, gin.H{
			"message": "Movie already have posters, please update, or delete to add posters",
		})
		return
	}
	err = m.DB.PosterRepo.AddPostersMovieById(movieIdNum, data.MainPoster, data.SecondPoster, data.ThirdPoster, data.FourthPoster, data.FifthPoster)
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
