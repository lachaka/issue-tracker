package handler

import (
	utils "issue-tracker/cmd/utils"
	"issue-tracker/pkg/models"
	"issue-tracker/pkg/models/cql"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	Register(*gin.Context)
	Login(*gin.Context)
	Logout(*gin.Context)
	GetUserById(*gin.Context)
}

type userHandler struct {
	userModel cql.UserModel
}

func NewUserHandler(userModel *cql.UserModel) UserHandler {
	return &userHandler{userModel: *userModel}
}

func (u *userHandler) Register(c *gin.Context) {
	var user models.User
	getUser, _ := c.Get("user")
	user = getUser.(models.User)

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		utils.Logger.ErrorLog.Println(err)
		c.JSON(http.StatusInternalServerError, "invalid password")
		return
	}

	user.Password = string(hash)

	_, err = u.userModel.Save(user)
	if err != nil {
		utils.Logger.ErrorLog.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	utils.Logger.InfoLog.Println("User created")
	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func (u *userHandler) Login(c *gin.Context) {

}

func (u *userHandler) Logout(c *gin.Context) {

}

func (u *userHandler) GetUserById(c *gin.Context) {

	id := c.Param("id")

	user, err := u.userModel.GetById(id)
	if err != nil {
		utils.Logger.ErrorLog.Print(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
