package routes

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// get all favorites
func (m *Manager) GET_Favorites(c *gin.Context) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	user, err := m.DB.GetUserById(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err.Error())
			c.JSON(404, gin.H{
				"message": "User not found",
			})
			return

		}

		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
	}

	favorites, err := m.DB.GetFavoritesByUserId(user.Id)
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

// add fav to user
func (m *Manager) POST_Favorite(c *gin.Context) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	id := c.Param("id")
	valid, movieId := utils.IsValidNum(id)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	err = m.DB.CheckUserExistsById(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err.Error())
			c.JSON(404, gin.H{
				"message": "User not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
	}

	err = m.DB.CheckMovieExistsById(movieId)
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
	added, err := m.DB.CheckIfMovieAdded(userId, movieId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	if added {
		log.Println("movie already added by this user")
		c.JSON(http.StatusConflict, gin.H{
			"message": "Movie already added to favorites by this user",
		})
		return
	}
	favorite, err := m.DB.AddToFavorites(userId, movieId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message":  "Added to favorites",
		"favorite": favorite,
	})
}

func (m *Manager) DELETE_Favorite(c *gin.Context) {
	token := c.GetHeader("Authorization")
	movieId := c.Param("id")

	valid, movieIdNum := utils.IsValidNum(movieId)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	favorites, err := m.DB.GetFavoritesByUserId(userId)
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
	exists := false
	favId := 0
	for _, v := range favorites {
		if v.MovieID == movieIdNum {
			exists = true
			favId = v.ID
			break
		}
	}
	if !exists {
		c.JSON(404, gin.H{
			"message": "No such movie added to favorites",
		})
		return
	}
	err = m.DB.DeleteFavoritesById(favId)
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

// deletes all favorites of CURRENT authenticated user
func (m *Manager) DELETE_Favorites(c *gin.Context) {
	token := c.GetHeader("Authorization")

	userId, err := utils.VerifyToken(token)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	err = m.DB.DeleteAllFavoritesByUserId(userId)
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
