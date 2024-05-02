package seasons

import (
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Tags Seasons
func (m *SeasonsRoute) DELETE_Season(c *gin.Context) {
	seasonId := c.Param("id")
	userRole := c.GetString("role")

	if userRole != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	valid, seasonIdNum := utils.IsValidNum(seasonId)
	if !valid {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	err := m.DB.SeasonRepository.DeleteSeason(seasonIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Season deleted",
	})
}

// @Tags Seasons
func (m *SeasonsRoute) DELETE_SeasonNumber(c *gin.Context) {
	movieId := c.Param("id")
	seasonNumber := c.Param("seasonNumber")
	userRole := c.GetString("role")

	if userRole != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	valid, movieIdNum := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	valid, seasonNumberNum := utils.IsValidNum(seasonNumber)
	if !valid {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	err := m.DB.SeasonRepository.DeleteSeasonNumberOfSelectedMovie(movieIdNum, seasonNumberNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "All Season numbers deleted",
	})
}
func (m *SeasonsRoute) DELETE_AllEpisodesOfSeason(c *gin.Context) {
	movieId := c.Param("id")
	userRole := c.GetString("role")

	if userRole != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	valid, movieIdNum := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	err := m.DB.SeasonRepository.DeleteAllSeasonsOfCurrentMovie(movieIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Seasons of movie deleted",
	})

}
