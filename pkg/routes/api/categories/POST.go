package categories

import (
	"log"
	"net/http"
	"project/internal/database/categories"

	"github.com/gin-gonic/gin"
)

// POST_Category creates a new category
// @Tags			categories
// @Summary		Create a category
// @Description	Creates a new category
// @Produce		json
// @Security		ApiKeyAuth
// @Param			category	body		routes.CategoryRequest			true	"Category object to be created"
// @Success		200			{object}	routes.DefaultMessageResponse	"Category Created"
// @Failure		400			{object}	routes.DefaultMessageResponse	"Bad request"
// @Failure		401			{object}	routes.DefaultMessageResponse	"Unauthorized"
// @Failure		409			{object}	routes.DefaultMessageResponse	"Category already exists"
// @Failure		500			{object}	routes.DefaultMessageResponse	"Internal server error"
// @Router			/api/categories [post]
func (m *CategoriesRoute) POST_Category(c *gin.Context) {
	userRole := c.GetString("role")
	userId := c.GetInt("userId")

	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	var category categories.Category

	err := c.BindJSON(&category)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	exists, err := m.DB.CategoriesRepository.CheckCategoryExistsByName(category.Name)
	if exists {
		c.JSON(409, gin.H{
			"message": "Category already exists",
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

	_, err = m.DB.CategoriesRepository.CreateCategory(userId, category.Name, category.Description)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Category Created",
	})

}
