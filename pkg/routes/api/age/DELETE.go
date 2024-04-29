package age

import (
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// DELETE_AgeCategory deletes an age category
//	@Summary		Delete an age category
//	@Description	Deletes an age category with the specified ID
//	@Produce		json
//	@Param			id	path	int	true	"Age Category ID"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	routes.DefaultMessageResponse	"Category Deleted"
//	@Failure		400	{object}	routes.DefaultMessageResponse	"Bad request"
//	@Failure		401	{object}	routes.DefaultMessageResponse	"Unauthorized"
//	@Failure		404	{object}	routes.DefaultMessageResponse	"Age category Not found"
//	@Failure		409	{object}	routes.DefaultMessageResponse	"Age category is Used in movies"
//	@Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
//	@Router			/api/age-categories/{id} [delete]
func (m *AgeRoute) DELETE_AgeCategory(c *gin.Context) {
	ageId := c.Param("id")
	userRole := c.GetString("role")
	userId := c.GetInt("userId")

	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	valid, ageIdNum := utils.IsValidNum(ageId)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	exists, err := m.DB.AgeRepository.CheckAgeCategoryExistsId(ageIdNum)
	if !exists {
		c.JSON(404, gin.H{
			"message": "Age category Not found",
		})
		return
	}

	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	used, err := m.DB.AgeRepository.CheckAgeCategoryIsUsedInMovies(ageIdNum)
	if used {
		c.JSON(409, gin.H{
			"message": "Age category is used in movies",
		})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	err = m.DB.AgeRepository.DeleteAgeCategoryById(ageIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "age Category Deleted",
	})
}
