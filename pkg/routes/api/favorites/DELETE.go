package favorites

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Tags			favorites
// @Summary		Deletes favorite movie by favorite ID
// @Description	Deletes favorite movie from auth. user's favorites
// @Accepts		json
// @Produce		json
// @Security		ApiKeyAuth
// @Param			id	path		int								true	"Movie id"
// @Success		200	{object}	routes.DefaultMessageResponse	"Favorite Deleted"
// @Failure		400	{object}	routes.DefaultMessageResponse	"Bad request"
// @Failure		401	{object}	routes.DefaultMessageResponse	"Unauthorised"
// @Failure		404	{object}	routes.DefaultMessageResponse	"No such movie added to favorites"
// @Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
// @Router			/api/favorites/{id} [DELETE]
func (m *FavoritesRoute) DELETE_Favorite(c *gin.Context) {
	movieId := c.Param("id")
	userId := c.GetInt("userId")
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorised",
		})
		return
	}
	valid, movieIdNum := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}

	favorites, err := m.DB.FavoritesRepository.GetFavoritesByUserId(userId) // get all favorites to check if this favorite movie really added to fav
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "No favorites added",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	favId := 0
	for _, v := range favorites {
		if v.MovieID == movieIdNum {
			favId = v.ID
			break
		}
	}
	if favId == 0 {
		c.JSON(404, gin.H{
			"message": "No such movie added to favorites",
		})
		return
	}
	err = m.DB.FavoritesRepository.DeleteFavoritesByMovieId(userId, movieIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Favorite Deleted",
	})

}

// @Tags			favorites
// @Summary		Deletes all favorite movies
// @Description	Deletes all favorite movies from auth. user's favorites
// @Accepts		json
// @Produce		json
// @Security		ApiKeyAuth
// @Success		200	{object}	routes.DefaultMessageResponse	"Favorites Cleared"
// @Failure		401	{object}	routes.DefaultMessageResponse	"Unauthorised"
// @Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
// @Router			/api/favorites/clear [DELETE]
func (m *FavoritesRoute) DELETE_Favorites(c *gin.Context) {
	userId := c.GetInt("userId")
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorised",
		})
		return
	}
	err := m.DB.FavoritesRepository.DeleteAllFavoritesByUserId(userId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Favorites cleared",
	})
}
