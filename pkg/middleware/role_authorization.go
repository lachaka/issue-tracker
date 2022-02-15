package middleware

import (
	"issue-tracker/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateUserRole(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		getUser := c.MustGet("user")
		user = *getUser.(*models.User)

		for _, r := range roles {
			if r == user.Role {
				c.Next()
				return
			}
		}

		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
