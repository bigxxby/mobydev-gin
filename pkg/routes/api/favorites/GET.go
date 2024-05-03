package favorites

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags			favorites
// @Summary		Get favorite movies
// @Description	Get all favorite movies of an auth. user
// @Produce		json
// @Security		ApiKeyAuth
// @Success		200	{object}	routes.DefaultMessageResponse	"Favorites"
// @Failure		401	{object}	routes.DefaultMessageResponse	"Unauthorised"
// @Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
// @Router			/api/favorites/ [get]
func (m *FavoritesRoute) GET_Favorites(c *gin.Context) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	userId := c.GetInt("userId")
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorised",
		})
		return
	}
	favorites, err := m.DB.FavoritesRepository.GetFavoritesByUserId(userId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	if favorites == nil {
		c.JSON(404, gin.H{
			"message": "No favorite movies added",
		})
		return
	}
	c.JSON(200, favorites)
}
