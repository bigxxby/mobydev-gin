package routes

import (
	"log"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

func (m *Manager) GET_Project(c *gin.Context) {
	limit := c.Query("limit")
	if limit != "" {
		valid, num := utils.IsValidNum(limit)
		if !valid {
			log.Println("number not valid")
			c.JSON(400, gin.H{
				"message": "Bad request",
			})
			return
		}
		projects, err := m.DB.GetProjects(num)
		if err != nil {
			log.Println(err.Error())
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}
		c.JSON(200, gin.H{
			"projects": projects,
		})
		return
	}

	// Если limit не указан, получить все проекты
	projects, err := m.DB.GetProjects(0) // 0 или другое значение по умолчанию
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"projects": projects,
	})
}

func (m *Manager) GET_ProjectById(c *gin.Context) {
	projectId := c.Param("id")
	valid, num := utils.IsValidNum(projectId)
	if !valid {
		log.Println("num is not valid")
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	project, err := m.DB.GetProjectById(num)
	if project == nil {
		c.JSON(404, gin.H{
			"message": "Project not found",
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
	c.JSON(200, project)
}
