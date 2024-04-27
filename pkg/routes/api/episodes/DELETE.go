package episodes

import (
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

func (m *EpisodesRoute) DELETE_Episode(c *gin.Context) {
	episodeId := c.Param("id")
	userRole := c.GetString("role")
	if userRole != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	valid, episodeIdNum := utils.IsValidNum(episodeId)
	if !valid {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	err := m.DB.EpisodeRepository.DeleteEpisodeById(episodeIdNum)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Episode deleted",
	})

}
func (m *EpisodesRoute) DELETE_AllEpisodesOfSeason(c *gin.Context) {

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
	err := m.DB.EpisodeRepository.DeleteAllEpisodesByIdOfSeason(seasonIdNum)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "All episodes of selected season deleted",
	})
}
func (m *EpisodesRoute) DELETE_EpisodeOfCurrentSeason(c *gin.Context) {
	seasonId := c.Param("id")
	episodeNumber := c.Param("episodeNumber")

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
	valid, episodeNumberNum := utils.IsValidNum(episodeNumber)
	if !valid {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	err := m.DB.EpisodeRepository.DeleteEpisodeByNumberOfSelectedSeason(seasonIdNum, episodeNumberNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Episode numbers of selected season deleted",
	})
}
