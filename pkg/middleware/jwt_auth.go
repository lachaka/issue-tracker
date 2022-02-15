package middleware

import (
	"issue-tracker/cmd/utils"
	"issue-tracker/pkg/models/cql"
	"issue-tracker/pkg/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

func Authorization(session *gocql.Session) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.JWTAuthService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)

			user, err := cql.GetByEmail(claims["email"].(string), session)
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			c.Set("user", user)
			c.Next()
		} else {
			utils.Logger.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
