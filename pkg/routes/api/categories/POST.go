package categories

import (
	"log"
	"net/http"
	"project/internal/database/categories"

	"github.com/gin-gonic/gin"
)

//	type Category struct {
//		ID          int    `json:"id"`
//		UserID      int    `json:"user_id"` // created by
//		Name        string `json:"category_name"`
//		Description string `json:"description"`
//		Created_at  string `json:"created_at"`
//	}
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
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Internal server error",
		})
		return
	}

	_, err = m.DB.CategoriesRepository.CreateCategory(userId, category.Name, category.Description)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Category Created",
	})

}
