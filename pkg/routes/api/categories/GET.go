package categories

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// GET_Category retrieves a category
//	@Summary		Get a category
//	@Description	Retrieves the category with the specified ID
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id	path		int								true	"Category ID"
//	@Success		200	{object}	routes.DefaultMessageResponse	"Category Retrieved"
//	@Failure		400	{object}	routes.DefaultMessageResponse	"Bad request"
//	@Failure		401	{object}	routes.DefaultMessageResponse	"Unauthorized"
//	@Failure		404	{object}	routes.DefaultMessageResponse	"Category not found"
//	@Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
//	@Router			/api/categories/{id} [get]
func (m *CategoriesRoute) GET_Category(c *gin.Context) {
	categoryId := c.Param("id")
	valid, categoryIdNums := utils.IsValidNum(categoryId)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}

	category, err := m.DB.CategoriesRepository.GetCategoryById(categoryIdNums)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Category not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, category)

}

// GET_Categories retrieves all categories
//	@Summary		Get all categories
//	@Description	Retrieves all categories
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{object}	routes.DefaultMessageResponse	"Categories Retrieved"
//	@Failure		401	{object}	routes.DefaultMessageResponse	"Unauthorized"
//	@Failure		404	{object}	routes.DefaultMessageResponse	"Categories not found"
//	@Failure		500	{object}	routes.DefaultMessageResponse	"Internal server error"
//	@Router			/api/categories [get]
func (m *CategoriesRoute) GET_Categories(c *gin.Context) {
	category, err := m.DB.CategoriesRepository.GetCategories()
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Category not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, category)

}
