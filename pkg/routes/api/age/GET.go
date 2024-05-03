package age

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Tags			ageCategory
// @Summary		Get all age categories
// @Description	Retrieves all age categories
// @Produce		json
// @Success		200	{object}	[]AgeCategory					"OK"
// @Failure		404	{objects}	routes.DefaultMessageResponse	"Age categories not found"
// @Failure		500	{object}	routes.DefaultMessageResponse	"Internal Server Error"
// @Router			/api/age-categories [get]
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

// @Tags			ageCategory
// @Summary		Get an age category by ID
// @Description	Retrieves an age category with the specified ID
// @Produce		json
// @Param			id	path		int								true	"Age Category ID"
// @Success		200	{object}	AgeCategory						"OK"
// @Failure		400	{object}	routes.DefaultMessageResponse	"Bad request"
// @Failure		404	{object}	routes.DefaultMessageResponse	"Age category not found"
// @Failure		500	{object}	routes.DefaultMessageResponse	"Internal Server Error"
// @Router			/api/age-categories/{id} [get]
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
