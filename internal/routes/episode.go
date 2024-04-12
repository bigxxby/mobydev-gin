package routes

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

func (m *Manager) GET_Episode(c *gin.Context) {
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
