package favorites

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

//	@Tags			favorites
//	@Summary		Adds a movie to favorites
//	@Description	Adds a movie with the specified ID to the favorites list of the authenticated user
//	@Accepts		json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id	path		int								true	"ID of the movie to add to favorites"
//	@Success		200	{object}	routes.DefaultMessageResponse	"Movie added to favorites"
//	@Failure		400	{object}	routes.DefaultMessageResponse	"Bad request"
//	@Failure		401	{object}	routes.DefaultMessageResponse	"Unauthorised"
//	@Failure		404	{object}	routes.DefaultMessageResponse	"Movie not found"
//	@Failure		409	{object}	routes.DefaultMessageResponse	"Movie already added to favorites by this user"
//	@Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
//	@Router			/api/favorites/{id} [POST]
func (m *FavoritesRoute) POST_Favorite(c *gin.Context) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	id := c.Param("id")
	valid, movieId := utils.IsValidNum(id)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	userId := c.GetInt("userId")
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorised",
		})
		return
	}

	err := m.DB.MovieRepository.CheckMovieExistsById(movieId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Movie not found",
			})
			return
		}
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	added, err := m.DB.FavoritesRepository.CheckIfMovieAddedToFavorites(userId, movieId) //checks if movie already added to users favorites
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	if added { // return StatusConflict if added
		log.Println("movie already added by this user")
		c.JSON(http.StatusConflict, gin.H{
			"message": "Movie already added to favorites by this user",
		})
		return
	}
	favorite, err := m.DB.FavoritesRepository.AddToFavorites(userId, movieId) //else add to favorites
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Movie added to favorites",
		"movieId": favorite.MovieID,
	})
}
