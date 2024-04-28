package middleware

import (
	"log"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware return userId , userRole by auth header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			return
		}
		userId, userRole, err := utils.VerifyToken(token)
		if err != nil {
			log.Println(err.Error())
			return
		}
		c.Set("userId", userId)
		c.Set("role", userRole)

		c.Next()
	}
}
