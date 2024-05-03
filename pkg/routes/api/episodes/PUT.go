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
//	@Summary		Updates episode
//	@Description	Updates episode by its ID
//	@Produce		json
//	@Param			id	path	int	true	"Episode ID"
//	@Security		ApiKeyAuth
//	@Param			episode	body		routes.EpisodeRequest			true	"Episode data"
//	@Success		200		{object}	routes.DefaultMessageResponse	"Episode updated"
//	@Failure		400		{object}	routes.DefaultMessageResponse	"Bad request"
//	@Failure		404		{object}	routes.DefaultMessageResponse	"Episode not found"
//	@Failure		500		{object}	routes.DefaultMessageResponse	"Internal Server Error"
//	@Router			/api/episodes/{id} [put]
func (m *EpisodesRoute) PUT_Episode(c *gin.Context) {
	episodeId := c.Param("id")
	userId := c.GetInt("userId")
	userRole := c.GetString("role")
	valid, episodeIdNum := utils.IsValidNum(episodeId)
	if !valid {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	if userId == 0 || userRole != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	var episode episode.Episode
	err := c.BindJSON(&episode)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	err = m.DB.EpisodeRepository.CheckEpisodeExistsById(episodeIdNum)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Episode not found",
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

	err = m.DB.EpisodeRepository.UpdateEpisode(episodeIdNum, userId, episode.URL, episode.EpisodeNumber, episode.Name, episode.DurationMinutes, releaseDate, episode.Description)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Episode updated",
	})

}
