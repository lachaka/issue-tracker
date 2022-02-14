package middleware

import (
	"fmt"
	"issue-tracker/cmd/utils"
	"issue-tracker/pkg/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.JWTAuthService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
			c.Next()
		} else {
			utils.Logger.ErrorLog.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
