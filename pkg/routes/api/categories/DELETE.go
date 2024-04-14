package categories

import (
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// delete category

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
		c.JSON(http.StatusUnauthorized, gin.H{
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
