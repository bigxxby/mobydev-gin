package posters

import (
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

func (m *PosterRoute) DELETE_Posters(c *gin.Context) {
	posterId := c.Param("id")
	userId := c.GetInt("userId")
	userRole := c.GetString("role")
	if userId == 0 || userRole != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"messsage": "Unauthorized",
		})
		return
	}
	valid, posterIdNum := utils.IsValidNum(posterId)
	if !valid {
		c.JSON(400, gin.H{
			"messsage": "Bad request",
		})
		return
	}
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
		})
		return
	}

	exists, err := m.DB.PosterRepo.CheckIfPosterExists(posterIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}
	if !exists {
		c.JSON(404, gin.H{
			"message": "Poster not found",
		})
		return
	}

	err = m.DB.PosterRepo.DeletePosters(posterIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Posters Deleted",
	})
}
func (m *PosterRoute) DELETE_PostersMovie(c *gin.Context) {
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
		c.JSON(400, gin.H{
			"messsage": "Bad request",
		})
		return
	}
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
		})
		return
	}

	err := m.DB.PosterRepo.DeletePostersOfMovie(movieIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Posters of movie Deleted",
	})
}
