package middleware

import (
	"issue-tracker/cmd/utils"
	"issue-tracker/pkg/models"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

func respondWithError(c *gin.Context, statusCode int, message interface{}) {
	c.AbortWithStatusJSON(statusCode, gin.H{"error": message})
}

func ValidateUserData() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		c.BindJSON(&user)

		_, err := mail.ParseAddress(user.Email)

		if err != nil {
			utils.Logger.ErrorLog.Println(err)
			respondWithError(c, http.StatusExpectationFailed, "Invalid email")
			return
		}

		if len(user.Password) < 8 {
			utils.Logger.ErrorLog.Println("password must be at least 8 symbols")
			respondWithError(c, http.StatusLengthRequired, "Password must be at least 8 symbols")
			return
		}

		user.Id = gocql.TimeUUID()

		c.Set("user", user)
		c.Next()
	}
}
