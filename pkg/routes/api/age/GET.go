package age

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

func (m *AgeRoute) GET_AgeCategories(c *gin.Context) {
	ageCategories, err := m.DB.AgeRepository.GetAllAgeCategories()
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Age categories not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, ageCategories)
}
func (m *AgeRoute) GET_AgeCategory(c *gin.Context) {
	ageId := c.Param("id")

	valid, ageIdNum := utils.IsValidNum(ageId)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	ageCategory, err := m.DB.AgeRepository.GetAgeCategoryById(ageIdNum)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Age category not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, ageCategory)
}
