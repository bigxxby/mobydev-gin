package posters

import (
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Summary		Deletes a poster
// @Description	Delete a poster by its ID
// @Tags			Posters
// @Produce		json
// @Param			id	path	string	true	"Poster ID"
// @Security		ApiKeyAuth
// @Success		200	{object}	routes.DefaultMessageResponse	"message": "Posters Deleted"
// @Failure		400	{object}	routes.DefaultMessageResponse	"message": "Bad request"
// @Failure		401	{object}	routes.DefaultMessageResponse	"message": "Unauthorized"
// @Failure		404	{object}	routes.DefaultMessageResponse	"message": "Poster not found"
// @Failure		500	{object}	routes.DefaultMessageResponse	"message": "Internal server error"
// @Router			/api/posters/{id}   [delete]
func (m *PosterRoute) DELETE_Posters(c *gin.Context) {
	posterId := c.Param("id")
	userId := c.GetInt("userId")
	userRole := c.GetString("role")
	if userId == 0 || userRole != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"messsage": "Unauthorized",
		})
		return
	}
	valid, posterIdNum := utils.IsValidNum(posterId)
	if !valid {
		c.JSON(400, gin.H{
			"messsage": "Bad request",
		})
		return
	}
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
		})
		return
	}

	exists, err := m.DB.PosterRepo.CheckIfPosterExists(posterIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}
	if !exists {
		c.JSON(404, gin.H{
			"message": "Poster not found",
		})
		return
	}

	err = m.DB.PosterRepo.DeletePosters(posterIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Posters Deleted",
	})
}

// @Tags	Posters
// @Summary		Deletes a poster of a movie
// @Description	Delete a poster by movie ID
// @Tags			posters
// @Produce		json
// @Param			id	path	string	true	"Movie ID"
// @Security		ApiKeyAuth
// @Success		200	{object}	routes.DefaultMessageResponse	"message": "Posters of a Movie Deleted"
// @Failure		400	{object}	routes.DefaultMessageResponse	"message": "Bad request"
// @Failure		401	{object}	routes.DefaultMessageResponse	"message": "Unauthorized"
// @Failure		500	{object}	routes.DefaultMessageResponse	"message": "Internal server error"
// @Router			/api/posters/movie/{id}   [delete]
func (m *PosterRoute) DELETE_PostersMovie(c *gin.Context) {
	movieId := c.Param("id")
	userId := c.GetInt("userId")
	userRole := c.GetString("role")
	if userId == 0 || userRole != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"messsage": "Unauthorized",
		})
		return
	}
	valid, movieIdNum := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(400, gin.H{
			"messsage": "Bad request",
		})
		return
	}

	err := m.DB.PosterRepo.DeletePostersOfMovie(movieIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Posters of movie Deleted",
	})
}
