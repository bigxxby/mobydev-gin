package episodes

import (
	"log"
	"net/http"
	"project/internal/database/episode"
	"project/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *EpisodesRoute) POST_Episode(c *gin.Context) {
	seasonId := c.Param("id")
	userId := c.GetInt("userId")
	userRole := c.GetString("role")
	if userRole != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	valid, seasonIdNums := utils.IsValidNum(seasonId)
	if !valid {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	err := m.DB.SeasonRepository.CheckSeasonExistsById(seasonIdNums)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Season not found",
		})
		return
	}
	var episode episode.Episode
	err = c.BindJSON(&episode)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	releaseDate, err := time.Parse("2006-01-02", episode.ReleaseDate)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	episodeId, err := m.DB.EpisodeRepository.CreateEpisode(userId, seasonIdNums, episode.EpisodeNumber, episode.URL, episode.Name, episode.Description, episode.DurationMinutes, releaseDate)
	if err != nil {
		log.Println(err.Error())

		return
	}
	c.JSON(200, gin.H{
		"message":   "Episode added",
		"episodeId": episodeId,
	})

}
func (m *EpisodesRoute) POST_Episodes(c *gin.Context) {
	seasonId := c.Param("id")
	userId := c.GetInt("userId")
	userRole := c.GetString("role")
	if userRole != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	valid, seasonIdNums := utils.IsValidNum(seasonId)
	if !valid {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	err := m.DB.SeasonRepository.CheckSeasonExistsById(seasonIdNums)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Season not found",
		})
		return
	}
	episodes := struct {
		Episodes []episode.Episode `json:"episodes"`
	}{}
	err = c.BindJSON(&episodes)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	if len(episodes.Episodes) == 0 {
		c.JSON(400, gin.H{
			"message": "At least one episode required",
		})
		return
	}
	episodeIds := []int{}
	for _, episode := range episodes.Episodes {

		releaseDate, err := time.Parse("2006-01-02", episode.ReleaseDate)
		if err != nil {
			log.Println(err.Error())
			c.JSON(400, gin.H{
				"message": "Bad request",
			})
			return
		}

		episodeId, err := m.DB.EpisodeRepository.CreateEpisode(userId, seasonIdNums, episode.EpisodeNumber, episode.URL, episode.Name, episode.Description, episode.DurationMinutes, releaseDate)
		if err != nil {
			log.Println(err.Error())

			return
		}
		episodeIds = append(episodeIds, episodeId)
	}
	c.JSON(200, gin.H{
		"message":    "Multiple Episodes added",
		"episodeIds": episodeIds,
	})

}
