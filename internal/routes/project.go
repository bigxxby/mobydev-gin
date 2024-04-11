package routes

import (
	"database/sql"
	"log"
	"net/http"
	"project/internal/database"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

func (m *Manager) GET_HTML_Project(c *gin.Context) {
	// token := c.GetHeader("Authorization")

	// userId, err := utils.VerifyToken(token)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"message": "Unauthorized",
	// 	})
	// 	return
	// }
	// user, err := m.DB.GetUserById(userId)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	c.JSON(500, gin.H{
	// 		"message": "Internal server error",
	// 	})
	// }
	// if user.Role != "admin" {
	// 	log.Println("this user is not admin")
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"message": "Unauthorized",
	// 	})
	// 	return
	// }
	c.HTML(200, "project_create.html", nil)
}

// get projects
func (m *Manager) GET_Projects(c *gin.Context) {
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

	projects, err := m.DB.GetProjects(0)
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

// get project by id
func (m *Manager) GET_Project(c *gin.Context) {
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
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Project not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, project)
}

// create project
func (m *Manager) POST_Project(c *gin.Context) {
	token := c.GetHeader("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user, err := m.DB.GetUserById(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "User not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return

	}

	if user.Role != "admin" {
		log.Println("this user is not admin")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	var project database.Project

	err = c.BindJSON(&project)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	_, err = m.DB.CreateProject(
		userId, project.
			ImageUrl,
		project.Name,
		project.Category,
		project.ProjectType,
		project.Year,
		project.AgeCategory,
		project.DurationMinutes,
		project.Keywords,
		project.Description,
		project.Director,
		project.Producer)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Project Created",
	})

}

// delete project
func (m *Manager) DELETE_Project(c *gin.Context) {
	token := c.GetHeader("Authorization")
	projectId := c.Param("id")

	valid, num := utils.IsValidNum(projectId)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user, err := m.DB.GetUserById(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "User not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return

	}
	if user.Role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	exists, err := m.DB.CheckProjectExistsById(num)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	if !exists {
		log.Println("project not found")
		c.JSON(404, gin.H{
			"message": "Project not found",
		})
		return
	}
	err = m.DB.DeleteProject(projectId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Project Deleted",
	})

}

// update project
func (m *Manager) PUT_Project(c *gin.Context) {
	token := c.GetHeader("Authorization")
	projectId := c.Param("id")

	valid, num := utils.IsValidNum(projectId)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad request",
		})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	user, err := m.DB.GetUserById(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "User not found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return

	}
	if user.Role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	var project database.Project
	err = c.BindJSON(&project)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	exists, err := m.DB.CheckProjectExistsById(num)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	if !exists {
		log.Println("project not found")
		c.JSON(404, gin.H{
			"message": "Project not found",
		})
		return
	}

	err = m.DB.UpdateProject(
		num, //project id we want to update
		project.ImageUrl,
		project.Name,
		project.Category,
		project.ProjectType,
		project.Year,
		project.AgeCategory,
		project.DurationMinutes,
		project.Keywords,
		project.Description,
		project.Director,
		project.Producer,
	)
	// err = m.DB.DeleteProject(projectId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Project Updated",
	})

}
