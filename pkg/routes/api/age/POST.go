package age

import (
	"log"
	"net/http"
	"project/internal/database/age"

	"github.com/gin-gonic/gin"
)

// POST_AgeCategory creates a new age category
//	@Summary		Create a new age category
//	@Description	Creates a new age category
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			ageCategory	body		routes.AgeCategoryRequest		true	"Age Category"
//	@Success		200			{object}	routes.DefaultMessageResponse	"Age category created"
//	@Failure		400			{object}	routes.DefaultMessageResponse	"Bad request"
//	@Failure		401			{object}	routes.DefaultMessageResponse	"Unauthorized"
//	@Failure		409			{object}	routes.DefaultMessageResponse	"Age category already exists"
//	@Failure		500			{object}	routes.DefaultMessageResponse	"Internal server error"
//	@Router			/api/age-categories [post]
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
			"message": "Age Category is not valid",
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
