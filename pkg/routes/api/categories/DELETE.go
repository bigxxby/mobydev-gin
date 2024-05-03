package categories

import (
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Tags			categories
// @Summary		Delete a category
// @Description	Deletes a category with the specified ID
// @Produce		json
// @Security		ApiKeyAuth
// @Param			id	path		int								true	"Category ID"
// @Success		200	{object}	routes.DefaultMessageResponse	"Category Deleted"
// @Failure		400	{object}	routes.DefaultMessageResponse	"Bad request"
// @Failure		401	{object}	routes.DefaultMessageResponse	"Unauthorized"
// @Failure		404	{object}	routes.DefaultMessageResponse	"Category not found"
// @Failure		400	{object}	routes.DefaultMessageResponse	"Cannot delete category because it is used in movies"
// @Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
// @Router			/api/categories/{id} [delete]
func (m *CategoriesRoute) DELETE_Category(c *gin.Context) {
	categoryId := c.Param("id")
	userRole := c.GetString("role")
	userId := c.GetInt("userId")

	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	valid, categoryIdNum := utils.IsValidNum(categoryId)
	if !valid {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	exists, err := m.DB.CategoriesRepository.CheckCategoryExistsById(categoryIdNum)
	if !exists {
		c.JSON(404, gin.H{
			"message": "Category not found",
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
	used, err := m.DB.CategoriesRepository.CheckCategoryIsUsedInMovies(categoryIdNum)
	if used {
		c.JSON(409, gin.H{
			"message": "Cannot delete category because it is used in movies",
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

	err = m.DB.CategoriesRepository.DeleteCategoryById(categoryId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Category Deleted",
	})

}
