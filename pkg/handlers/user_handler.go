package handler

import (
	utils "issue-tracker/cmd/utils"
	"issue-tracker/pkg/models"
	"issue-tracker/pkg/models/cql"
	"issue-tracker/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	Register(*gin.Context)
	Login(*gin.Context)
	Logout(*gin.Context)
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

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Logger.ErrorLog.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "invalid password")
		return
	}

	user.Password = string(hash)

	_, err = u.userModel.Save(user)
	if err != nil {
		utils.Logger.ErrorLog.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.Logger.InfoLog.Println("User created")
	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func (u *userHandler) Login(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	userData, err := u.userModel.GetByEmail(user.Email)

	if err != nil {
		utils.Logger.ErrorLog.Println(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		utils.Logger.ErrorLog.Println(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	token, err := service.JWTAuthService().GenerateToken(*userData)

	if err != nil {
		utils.Logger.ErrorLog.Println(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func (u *userHandler) Logout(c *gin.Context) {
	c.Status(http.StatusOK)
}
