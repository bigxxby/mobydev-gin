package episodes

import (
	"log"
	"net/http"
	"project/internal/database/episode"
	"project/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

//	@Tags			episodes
// POST_Episode creates a single episode for a season
//
//	@Summary		Create episode
//	@Description	Creates a single episode for the specified season
//	@Produce		json
//	@Param			id	path	int	true	"Season ID"
//	@Security		ApiKeyAuth
//	@Param			episode	body		routes.EpisodeRequest			true	"Episode"
//	@Success		200		{object}	routes.DefaultMessageResponse	"Episode added"
//	@Failure		400		{object}	routes.DefaultMessageResponse	"Bad request"
//	@Failure		404		{object}	routes.DefaultMessageResponse	"Season not found"
//	@Failure		500		{object}	routes.DefaultMessageResponse	"Internal Server Error"
//	@Router			/api/episodes/{id} [post]
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
			"message": "Invalid date please use format `2006-01-02`",
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

type EpisodeRequestBody struct {
	Episodes []EpisodeRequest `json:"episodes"`
}

type EpisodeRequest struct {
	URL             string `json:"url" binding:"required"`
	EpisodeNumber   int    `json:"episode_number" binding:"required"`
	Name            string `json:"name" binding:"required"`
	DurationMinutes int    `json:"duration_minutes" binding:"required"`
	ReleaseDate     string `json:"release_date" binding:"required"`
	Description     string `json:"description"`
}

// POST_Episodes creates multiple episodes for a season
//
//	@Summary		Create episodes
//	@Description	Creates multiple episodes for the specified season
//	@Produce		json
//	@Param			id	path	int	true	"Season ID"
//	@Security		ApiKeyAuth
//	@Param			episodes	body		EpisodeRequestBody				true	"Episodes"
//	@Success		200			{object}	routes.DefaultMessageResponse	"Multiple Episodes added"
//	@Failure		400			{object}	routes.DefaultMessageResponse	"Bad request"
//	@Failure		404			{object}	routes.DefaultMessageResponse	"Season not found"
//	@Failure		500			{object}	routes.DefaultMessageResponse	"Internal Server Error"
//	@Router			/api/episodes/{id}/multiple [post]
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
	episodes := EpisodeRequestBody{}
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
				"message": "Invalid date please use format `2006-01-02`",
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
