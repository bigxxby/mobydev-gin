package posters

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"
	"project/internal/utils/mapping"

	"github.com/gin-gonic/gin"
)

func (m *PosterRoute) GET_PostersOfMoive(c *gin.Context) {
	movieId := c.Param("id")
	valid, movieIdNum := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
		})
		return
	}

	posters, err := m.DB.PosterRepo.GetPostersAllOfMovieById(movieIdNum)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Posters not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch posters",
		})
		return
	}

	c.JSON(http.StatusOK, mapping.TrimPoster(*posters))
}
