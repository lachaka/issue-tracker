package handler

import (
	"net/http"
	"issue-tracker/pkg/models"
	"issue-tracker/pkg/models/cql"

	"github.com/gocql/gocql"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	CreateUser(*gin.Context)
	GetUserById(*gin.Context)
}

type userHandler struct {
	userModel cql.UserModel
}

func NewUserHandler(userModel *cql.UserModel) UserHandler {
	return &userHandler{userModel: *userModel}
}

func (u *userHandler) CreateUser(c *gin.Context) {
	var user models.User

	c.BindJSON(&user)
	user.Id = gocql.TimeUUID()
	
	data, err := u.userModel.Save(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": data})
}

func (u *userHandler) GetUserById(c *gin.Context) {

	id := c.Param("id")

	user, err := u.userModel.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}