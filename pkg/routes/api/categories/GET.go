package categories

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

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
