package age

import (
	"log"
	"net/http"
	"project/internal/database/age"

	"github.com/gin-gonic/gin"
)

func (m *AgeRoute) POST_AgeCategory(c *gin.Context) {
	userRole := c.GetString("role")
	userId := c.GetInt("userId")

	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var ageCategory age.AgeCategory

	// if (ageCategory.MaxAge < ageCategory.MinAge) || (ageCategory.MaxAge < ageCategory.MinAge) || (ageCategory.MaxAge > 100) || (ageCategory.MinAge < 0) {
	// 	log.Println("Min and Max age bad request")
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "Bad request",
	// 	})
	// 	return
	// }

	err := c.BindJSON(&ageCategory)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	valid := isValidAgeCategory(ageCategory)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	exists, err := m.DB.AgeRepository.CheckAgeCategoryExistsByName(ageCategory.Name)
	if exists {
		c.JSON(409, gin.H{
			"message": "Age category already exists",
		})
		return
	}
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	_, err = m.DB.AgeRepository.CreateAgeCategory(userId, ageCategory.Name, ageCategory.Note, ageCategory.MinAge, ageCategory.MaxAge)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Age category created",
	})
}
