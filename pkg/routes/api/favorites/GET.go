package favorites

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// get all favorites
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
	c.JSON(200, favorites)
}
