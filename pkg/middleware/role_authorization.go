package middleware

import (
	"issue-tracker/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateUserRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		getUser := c.MustGet("user")
		user = *getUser.(*models.User)

		if user.Role != role {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
