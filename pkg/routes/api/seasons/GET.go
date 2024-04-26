package seasons

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

func (m *SeasonsRoute) GET_Season(c *gin.Context) {
	seasonId := c.Param("id")
	valid, seasonIdNum := utils.IsValidNum(seasonId)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}

	season, err := m.DB.SeasonRepository.GetSeasonById(seasonIdNum)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Season not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, season)

}

func (m *SeasonsRoute) GET_AllSeasonsOfMovie(c *gin.Context) {
	movieId := c.Param("id")
	valid, movieIdNum := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	seasons, err := m.DB.SeasonRepository.GetAllSeasonsOfMovieId(movieIdNum)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "No seasons found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, gin.H{
		"seasons": seasons,
	})
}
