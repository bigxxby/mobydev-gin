package episodes

import (
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// DELETE_Episode deletes episode by its ID.
// @Tags			episodes
// @Summary		Delete an episode
// @Description	Deletes an episode with the specified ID
// @Produce		json
// @Security		ApiKeyAuth
// @Param			id	path		int								true	"Episode ID"
// @Success		200	{object}	routes.DefaultMessageResponse	"Episode deleted"
// @Failure		400	{object}	routes.DefaultMessageResponse	"Bad request"
// @Failure		401	{object}	routes.DefaultMessageResponse	"Unauthorized"
// @Failure		404	{object}	routes.DefaultMessageResponse	"Episode Not Found"
// @Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
// @Router			/api/episodes/{id} [delete]
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
	err := m.DB.EpisodeRepository.CheckEpisodeExistsById(episodeIdNum)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Episode not found",
		})
		return
	}
	err = m.DB.EpisodeRepository.DeleteEpisodeById(episodeIdNum)
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

// DELETE_AllEpisodesOfSeason deletes all episodes of a specific season.
// @Tags			episodes
// @Summary		Delete all episodes of a season
// @Description	Deletes all episodes belonging to the specified season
// @Produce		json
// @Security		ApiKeyAuth
// @Param			id	path		int								true	"Season ID"
// @Success		200	{object}	routes.DefaultMessageResponse	"All episodes of selected season deleted"
// @Failure		400	{object}	routes.DefaultMessageResponse	"Bad request"
// @Failure		401	{object}	routes.DefaultMessageResponse	"Unauthorized"
// @Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
// @Router			/api/episodes/season/{id}/clear [delete]
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

// DELETE_EpisodeOfCurrentSeason deletes an episode of the current season.
// @Tags			episodes
// @Summary		Delete an episode of the current season
// @Description	Deletes an episode with the specified number from the current season. (NOTE: deletes ALL episodes of selected number, including duplicates)
// @Produce		json
// @Security		ApiKeyAuth
// @Param			id				path	int	true	"Season ID"
// @Param			episodeNumber	path	int	true	"Episode number"
// @Success		200	{object}	routes.DefaultMessageResponse	"Episode numbers of selected season deleted"
// @Failure		400	{object}	routes.DefaultMessageResponse	"Bad request"
// @Failure		401	{object}	routes.DefaultMessageResponse	"Unauthorized"
// @Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
// @Router			/api/episodes/season/{id}/{episodeNumber} [delete]
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
