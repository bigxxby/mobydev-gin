package age

import (
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

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
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Internal server error",
		})
		return
	}

	err = m.DB.AgeRepository.DeleteAgeCategoryById(ageIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Category Deleted",
	})
}
