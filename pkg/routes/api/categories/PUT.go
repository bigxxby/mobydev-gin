package categories

import (
	"log"
	"net/http"
	"project/internal/database/categories"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// PUT_Category updates a category
// @Summary Update a category
// @Description Updates an existing category
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Category ID"
// @Param category body routes.CategoryRequest true "Updated category object"
// @Success 200 {object} routes.DefaultMessageResponse "Category Updated"
// @Failure 400 {object} routes.DefaultMessageResponse "Bad request"
// @Failure 401 {object} routes.DefaultMessageResponse "Unauthorized"
// @Failure 404 {object} routes.DefaultMessageResponse "Category Not found"
// @Failure 409 {object} routes.DefaultMessageResponse "Category with this name already exists"
// @Failure 500 {object} routes.DefaultMessageResponse "Internal server error"
// @Router /api/categories/{id} [put]
func (m *CategoriesRoute) PUT_Category(c *gin.Context) {
	categoryId := c.Param("id")
	userRole := c.GetString("role")
	userId := c.GetInt("userId")

	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var category categories.Category

	valid, categoryIdNum := utils.IsValidNum(categoryId)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	err := c.BindJSON(&category)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	exists, err := m.DB.CategoriesRepository.CheckCategoryExistsById(categoryIdNum)
	if !exists {
		c.JSON(404, gin.H{
			"message": "Category Not found",
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

	existingCategory, err := m.DB.CategoriesRepository.GetCategoryById(categoryIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if existingCategory.Name != category.Name {
		exists, err = m.DB.CategoriesRepository.CheckCategoryExistsByName(category.Name)
		if exists {
			c.JSON(409, gin.H{
				"message": "Category with this name already exists",
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
	}

	_, err = m.DB.CategoriesRepository.UpdateCategory(categoryIdNum, category.Name, category.Description)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Category Updated",
	})
}
