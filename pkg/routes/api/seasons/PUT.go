package seasons

import (
	"log"
	"net/http"
	"project/internal/database/season"
	"project/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

//	@Tags			seasons
//	@Summary		Updates a season
//	@Description	Updates the details of a season with the specified ID
//	@Param			id	path	int	true	"Season ID"
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			season	body		routes.SeasonBodyRequest		true	"Updated season information"
//	@Success		200		{object}	routes.DefaultMessageResponse	"Season updated"
//	@Failure		400		{object}	routes.DefaultMessageResponse	"Bad request"
//	@Failure		401		{object}	routes.DefaultMessageResponse	"Unauthorized"
//	@Failure		404		{object}	routes.DefaultMessageResponse	"Season not found"
//	@Failure		500		{object}	routes.DefaultMessageResponse	"Internal server error"
//	@Router			/api/seasons/{id} [PUT]
func (m *SeasonsRoute) PUT_Season(c *gin.Context) {
	seasonId := c.Param("id")
	userId := c.GetInt("userId")
	userRole := c.GetString("role")

	if userId == 0 || userRole != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	var season season.Season
	err := c.BindJSON(&season)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
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
	err = m.DB.SeasonRepository.CheckSeasonExistsById(seasonIdNum)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Season not found",
		})
		return
	}
	releaseDate, err := time.Parse("2006-01-02", season.ReleaseDate)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Invalid date please use format `2006-01-02`",
		})
		return
	}

	err = m.DB.SeasonRepository.UpdateSeason(seasonIdNum, season.SeasonNumber, season.Name, season.Description, releaseDate)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Season updated",
	})

}
