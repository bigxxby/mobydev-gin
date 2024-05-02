package episodes

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

//	@Tags			episodes
// GET_Episode retrieves an episode by ID
//
//	@Summary		Retrieve an episode
//	@Description	Retrieves an episode with the specified ID
//	@Produce		json
//	@Param			id	path		int								true	"Episode ID"
//	@Success		200	{object}	episode.Episode					"Episode"
//	@Failure		400	{object}	routes.DefaultMessageResponse	"Bad request"
//	@Failure		404	{object}	routes.DefaultMessageResponse	"Episode not found"
//	@Failure		500	{object}	routes.DefaultMessageResponse	"Internal Server Error"
//	@Router			/api/episodes/{id} [get]
func (m *EpisodesRoute) GET_Episode(c *gin.Context) {
	episodeId := c.Param("id")
	valid, episodeIdNums := utils.IsValidNum(episodeId)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}

	episode, err := m.DB.EpisodeRepository.GetEpisodeById(episodeIdNums)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Episode not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, episode)

}
