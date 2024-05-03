package seasons

import (
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

//	@Tags			seasons
//	@Summary		Deletes a season
//	@Description	Deletes a season with the specified ID
//	@Param			id	path	int	true	"Season ID"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	routes.DefaultMessageResponse	"Season deleted"
//	@Failure		400	{object}	routes.DefaultMessageResponse	"Bad request"
//	@Failure		401	{object}	routes.DefaultMessageResponse	"Unauthorized"
//	@Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
//	@Router			/api/seasons/{id} [DELETE]
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

//	@Tags			seasons
//	@Summary		Deletes all episodes of a specific season
//	@Description	Deletes all episodes of the specified season of a movie
//	@Param			id				path	int	true	"Movie ID"
//	@Param			seasonNumber	path	int	true	"Season number"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	routes.DefaultMessageResponse	"All Season numbers deleted"
//	@Failure		400	{object}	routes.DefaultMessageResponse	"Bad request"
//	@Failure		401	{object}	routes.DefaultMessageResponse	"Unauthorized"
//	@Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
//	@Router			/api/seasons/movie/{id}/{seasonNumber} [DELETE]
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

//	@Tags			seasons
//	@Summary		Deletes all episodes seasons of a movie
//	@Description	Deletes all seasons of the specified movie
//	@Param			id	path	int	true	"Movie ID"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	routes.DefaultMessageResponse	"All Seasons of movie deleted"
//	@Failure		400	{object}	routes.DefaultMessageResponse	"Bad request"
//	@Failure		401	{object}	routes.DefaultMessageResponse	"Unauthorized"
//	@Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
//	@Router			/api/seasons/movie/{id}/clear [DELETE]
func (m *SeasonsRoute) DELETE_AllSeasonsOfMovie(c *gin.Context) {
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
		"message": "All Seasons of movie deleted",
	})

}
