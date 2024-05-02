package age

import (
	"log"
	"net/http"
	"project/internal/database/age"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Tags			ageCategory
// PUT_AgeCategory updates an age category
//
//	@Summary		Update an age category
//	@Description	Updates an age category with the specified ID
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id			path		int								true	"Age Category ID"
//	@Param			ageCategory	body		routes.AgeCategoryRequest		true	"Age Category"
//	@Success		200			{object}	routes.DefaultMessageResponse	"Category Updated"
//	@Failure		400			{object}	routes.DefaultMessageResponse	"Bad request"
//	@Failure		401			{object}	routes.DefaultMessageResponse	"Unauthorized"
//	@Failure		404			{object}	routes.DefaultMessageResponse	"Age category Not found"
//	@Failure		409			{object}	routes.DefaultMessageResponse	"Age Category with this name already exists"
//	@Failure		500			{object}	routes.DefaultMessageResponse	"Internal server error"
//	@Router			/api/age-categories/{id} [put]
func (m *AgeRoute) PUT_AgeCategory(c *gin.Context) {
	ageId := c.Param("id")
	userRole := c.GetString("role")
	userId := c.GetInt("userId")

	if userRole != "admin" || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var ageCategory age.AgeCategory

	valid, ageIdNum := utils.IsValidNum(ageId)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	err := c.BindJSON(&ageCategory)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	valid = utils.IsValidAgeCategory(ageCategory)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Age category is not valid",
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
	existingAgeCategory, err := m.DB.AgeRepository.GetAgeCategoryById(ageIdNum)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if existingAgeCategory.Name != ageCategory.Name {
		exists, err := m.DB.AgeRepository.CheckAgeCategoryExistsByName(ageCategory.Name)
		if exists {
			c.JSON(409, gin.H{
				"message": "Age Category with this name already exists",
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
	}

	err = m.DB.AgeRepository.UpdateAgeById(ageIdNum, ageCategory.Name, ageCategory.Note, ageCategory.MinAge, ageCategory.MaxAge)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Category Updated",
	})
}
