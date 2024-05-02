package seasons

import (
	"log"
	"net/http"
	"project/internal/database/season"
	"project/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// @Tags Seasons
func (m *SeasonsRoute) POST_CreateSeason(c *gin.Context) {
	movieId := c.Param("id")
	userId := c.GetInt("userId")
	userRole := c.GetString("role")

	if userId == 0 || userRole != "admin" {
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
	err := m.DB.MovieRepository.CheckMovieExistsById(movieIdNum)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Movie not found",
		})
		return
	}
	var season season.Season
	err = c.BindJSON(&season)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	releaseDate, err := time.Parse("2006-01-02", season.ReleaseDate)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	seasonId, err := m.DB.SeasonRepository.AddSeasonToTheMovie(movieIdNum, season.UserID, season.Name, season.SeasonNumber, season.Description, releaseDate)
	if err != nil {

		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message":  "Season Added",
		"seasonId": seasonId,
	})
}

// @Tags Seasons
func (m *SeasonsRoute) POST_CreateSeasons(c *gin.Context) {
	movieId := c.Param("id")
	userId := c.GetInt("userId")
	userRole := c.GetString("role")

	if userId == 0 || userRole != "admin" {
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
	err := m.DB.MovieRepository.CheckMovieExistsById(movieIdNum)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Movie not found",
		})
		return
	}
	seasons := struct {
		Seasons []season.Season `json:"seasons"`
	}{}
	err = c.BindJSON(&seasons)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	seasonIds := []int{}

	for _, season := range seasons.Seasons {

		releaseDate, err := time.Parse("2006-01-02", season.ReleaseDate)
		if err != nil {
			log.Println(err.Error())
			c.JSON(400, gin.H{
				"message": "Bad request",
			})
			return
		}

		seasonId, err := m.DB.SeasonRepository.AddSeasonToTheMovie(movieIdNum, season.UserID, season.Name, season.SeasonNumber, season.Description, releaseDate)
		if err != nil {

			log.Println(err.Error())
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}
		seasonIds = append(seasonIds, seasonId)
	}

	c.JSON(200, gin.H{
		"message":   "Season Added",
		"seasonIds": seasonIds,
	})
}
